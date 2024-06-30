package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/Brett-Tanner/go_with_tests/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: post 2")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
