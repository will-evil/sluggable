package sluggable

import (
	"errors"
	"fmt"
	"github.com/bradfitz/slice"
	slugpkg "github.com/gosimple/slug"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const (
	defCount = 2
	tagName  = "sluggable"
)

// UniqueFunc is func for check slug unique.
type UniqueFunc func(slug string) (bool, error)

// CreateSlug makes slug from obj.
// obj must implement Sluggabler interface.
func CreateSlug(obj Sluggabler) (string, error) {
	slugLang, err := obj.SlugLang()
	if err != nil {
		return "", err
	}

	return MakeSlug(obj, slugLang, obj.IsSlugUnique)
}

func MakeSlug(obj interface{}, language string, isUnique UniqueFunc) (string, error) {
	counter := defCount

	slugBricks, err := slugBricks(obj)
	if err != nil {
		return "", err
	} else if len(slugBricks) == 0 {
		return "", errors.New("slug bricks can not be empty")
	}
	lastIndex := len(slugBricks)

	for {
		slug := slugpkg.MakeLang(strings.Join(slugBricks, "-"), language)

		if isUnique, err := isUnique(slug); err != nil {
			return "", nil
		} else if isUnique {
			return slug, nil
		} else if counter == defCount {
			slugBricks = append(slugBricks, string(counter))
			continue
		}

		counter++
		slugBricks[lastIndex] = string(counter)
	}
}

func slugBricks(obj interface{}) ([]string, error) {
	type Brick struct {
		index uint64
		value string
	}
	var elements []*Brick
	var result []string

	t := reflect.TypeOf(obj)
	tv := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get(tagName)
		if tag == "" {
			continue
		}

		filedValue, err := valToStr(tv.Field(i), field.Type.Kind())
		if err != nil {
			return result, err
		}

		slugIndex, err := strconv.ParseUint(tag, 10, 64)
		if err != nil {
			return result, fmt.Errorf("convert '%s' value '%s' to uint is failed", tagName, tag)
		}
		elements = append(elements, &Brick{index: slugIndex, value: filedValue})
	}

	sf := slice.SortInterface(elements[:], func(i, j int) bool {
		return elements[i].index < elements[j].index
	})
	sort.Sort(sf)

	for _, brick := range elements {
		result = append(result, brick.value)
	}

	return result, nil
}

func valToStr(value reflect.Value, kind reflect.Kind) (string, error) {
	switch kind {
	case reflect.String:
		return value.String(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10), nil
	default:
		return "", fmt.Errorf("%s is not support type", kind.String())
	}
}
