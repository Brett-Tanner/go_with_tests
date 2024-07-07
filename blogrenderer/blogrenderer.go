package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type PostRenderer struct {
	templ *template.Template
}

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func (p Post) SanitizedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "_", -1))
}

//go:embed "templates/*"
var postTemplates embed.FS

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	post.Body = string(mdToHtml([]byte(post.Body)))

	return r.templ.ExecuteTemplate(w, "blog.gohtml", post)
}

func mdToHtml(md []byte) template.HTML {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	htmlRenderer := html.NewRenderer(opts)

	return template.HTML(markdown.Render(doc, htmlRenderer))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}
