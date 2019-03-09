# Sluggable
Package for creating slugs from structure fields.

## Installation
````bash
go get github.com/will-evil/sluggable
````

## Usage
````
import (
    "github.com/will-evil/sluggable"
    fmt
)

type Post struct {
    Title string `sluggable:"1"`
}

func (p Post) IsSlugUnique() (bool, error) {
    return isSlugExists(slug), nil
}

func (p Post) SlugLang() (string, error) {
	return "en", nil
}

func main() {
    post := Post{Title: "My title"}
    slug, err := sluggable.CreateSlug(post)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(slug)
    // my-title
}
````