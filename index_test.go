package sluggable

import (
	"reflect"
	"testing"
)

type sluggable struct {
}

func (s sluggable) IsSlugUnique(slug string) (bool, error) {
	return true, nil
}

func (s sluggable) SlugLang() (string, error) {
	return "en", nil
}

func TestCreateSlug(t *testing.T) {
	type Post struct {
		sluggable
		Title string `sluggable:"1"`
	}

	testList := []struct {
		obj Post
		res string
	}{
		{
			obj: Post{Title: "dogs & cats"},
			res: "dogs-and-cats",
		},
		{
			obj: Post{Title: "123 my number ! title ()"},
			res: "123-my-number-title",
		},
		{
			obj: Post{Title: "  123456  "},
			res: "123456",
		},
		{
			obj: Post{Title: "*&(&---@"},
			res: "and-and-at",
		},
	}

	for _, testSet := range testList {
		slug, err := CreateSlug(testSet.obj)
		if err != nil {
			t.Fatal(err)
		}
		if testSet.res != slug {
			t.Errorf("Expected: %s, got: %s", testSet.res, slug)
		}
	}
}

func TestCreateSlug2(t *testing.T) {
	type Post struct {
		sluggable
		Title string `sluggable:"1"`
		Year  uint   `sluggable:"2"`
	}

	testList := []struct {
		obj Post
		res string
	}{
		{
			obj: Post{Title: "Created in", Year: 1999},
			res: "created-in-1999",
		},
	}

	for _, testSet := range testList {
		slug, err := CreateSlug(testSet.obj)
		if err != nil {
			t.Fatal(err)
		}
		if testSet.res != slug {
			t.Fatalf("Expected: %s, got: %s", testSet.res, slug)
		}
	}
}

func TestCreateSlug3(t *testing.T) {
	type Post struct {
		sluggable
		Year  uint   `sluggable:"2"`
		Title string `sluggable:"1"`
	}

	testList := []struct {
		obj Post
		res string
	}{
		{
			obj: Post{Title: "Created in", Year: 1999},
			res: "created-in-1999",
		},
	}

	for _, testSet := range testList {
		slug, err := CreateSlug(testSet.obj)
		if err != nil {
			t.Fatal(err)
		}
		if testSet.res != slug {
			t.Fatalf("Expected: %s, got: %s", testSet.res, slug)
		}
	}
}

func TestValToStr(t *testing.T) {
	var (
		valUint   uint   = 4294967295
		valUint8  uint8  = 255
		valUint16 uint16 = 65535
		valUint32 uint32 = 4294967295
		valUint64 uint64 = 5294967295
	)

	var (
		valInt         = 65535
		valInt8  int8  = 127
		valInt16 int16 = 32767
		valInt32 int32 = 2147483647
		valInt64 int64 = 5294967295
	)

	testList := []struct {
		val interface{}
		res string
	}{
		{
			val: valUint,
			res: "4294967295",
		},
		{
			val: valUint8,
			res: "255",
		},
		{
			val: valUint16,
			res: "65535",
		},
		{
			val: valUint32,
			res: "4294967295",
		},
		{
			val: valUint64,
			res: "5294967295",
		},
		{
			val: valInt,
			res: "65535",
		},
		{
			val: valInt8,
			res: "127",
		},
		{
			val: valInt16,
			res: "32767",
		},
		{
			val: valInt32,
			res: "2147483647",
		},
		{
			val: valInt64,
			res: "5294967295",
		},
	}

	for _, testSet := range testList {
		value := reflect.ValueOf(testSet.val)
		valType := reflect.TypeOf(testSet.val).Kind()

		strVal, err := valToStr(value, valType)
		if err != nil {
			t.Fatal(err)
		}
		if testSet.res != strVal {
			t.Fatalf("Expected: %s, got: %s", testSet.res, strVal)
		}
	}
}
