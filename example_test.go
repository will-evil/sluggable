package sluggable_test

import (
	"fmt"
	"sluggable"
)

// Post hast title to represent a post.
type Post struct {
	Title string `sluggable:"1"`
}

// IsSlugUnique check slug uniqueness.
func (p Post) IsSlugUnique(slug string) (bool, error) {
	return isSlugExists(slug), nil
}

// SlugLang provide lang for creating slug.
func (p Post) SlugLang() (string, error) {
	return "en", nil
}

// ExampleCreateSlug create slug from post.
func ExampleCreateSlug() {
	post := Post{Title: "dogs & cats"}

	slug, err := sluggable.CreateSlug(post)
	if err != nil {
		panic(err)
	}

	// dogs-and-cats
	fmt.Println(slug)
}

func isSlugExists(slug string) bool {
	existsSlugs := []string{
		"slug-1",
		"slug-2",
	}

	for _, existsSlug := range existsSlugs {
		if slug == existsSlug {
			return false
		}
	}

	return true
}
