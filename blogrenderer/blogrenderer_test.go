package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"

	"github.com/Brett-Tanner/go_with_tests/blogrenderer"
)

var post = blogrenderer.Post{
	Title: "hello world",
	Body: `# This is a h1
[here's a link](https://google.com)`,
	Description: "This is a description",
	Tags:        []string{"go", "tdd"},
}

func TestRender(t *testing.T) {
	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, post); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello 1"}, {Title: "Hello 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/posts/hello_1">Hello 1</a></li><li><a href="/posts/hello_2">Hello 2</a></li></ol>`

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, post)
	}
}
