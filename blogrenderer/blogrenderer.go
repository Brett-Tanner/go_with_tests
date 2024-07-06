package blogrenderer

import (
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, post Post) error {
	_, err := fmt.Fprintf(w, `<h1>%s</h1>
<p>%s</p>
Tags: %s`, post.Title, post.Description, arrayToUList(post.Tags))

	return err
}

func arrayToUList(listItems []string) string {
	list := strings.Builder{}
	list.WriteString("<ul>")

	for _, item := range listItems {
		fmt.Fprintf(&list, "<li>%s</li>", item)
	}

	list.WriteString("</ul>")

	return list.String()
}
