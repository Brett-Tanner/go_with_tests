package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/Brett-Tanner/go_with_tests/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title:Post 1
Description:Description 1
Tags:Python,< Ruby
---
This is an essay on why Ruby is the greatest
And Go as well I guess`
		secondBody = `Title:Post 2
Description:Description 2
Tags:Tag 2`
	)

	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"Python", "< Ruby"},
		Body: `This is an essay on why Ruby is the greatest
And Go as well I guess`,
	}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
