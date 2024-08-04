package fs

import (
	"embed"
)

//go:embed templates/*.tmpl
var templates embed.FS

func NewTemplates() *embed.FS {
	return &templates
}
