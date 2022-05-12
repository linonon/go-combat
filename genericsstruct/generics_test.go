package generics

import (
	"fmt"
	"testing"
)

func TestGenerics(t *testing.T) {
	category := Category{
		ID:   1,
		Name: "Go Generics",
		Slug: "go-generics",
	}

	// create cache for blog.Category struct
	categoryCache := New[Category]()
	// add category to cache
	categoryCache.Set(category.Slug, category)
	fmt.Printf("cp get:%+v\n", categoryCache.Get(category.Slug))

	post := Post{
		ID: 1,
		Categories: []Category{
			{ID: 1, Name: "Go Generics", Slug: "go-generics"},
		},
		Title: "Generics in Golang structs",
		Text:  "Here go's the text",
		Slug:  "generics-in-golang-structs",
	}
	// create cache for blog.Post struct
	postCache := New[Post]()
	postCache.Set(post.Slug, post)
	fmt.Printf("cp get:%+v\n", postCache.Get(post.Slug))
}
