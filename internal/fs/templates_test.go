package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderTemplate(t *testing.T) {
	tests := []struct {
		name     string
		template string
	}{
		{
			name:     "Markdown Template",
			template: "templates/markdown.tmpl",
		},
		{
			name:     "HTML Template",
			template: "templates/html.tmpl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tpl := NewTemplates()
			bb, err := tpl.ReadFile(tt.template)

			assert.NoError(t, err)
			assert.NotNil(t, bb)

			// Add additional assertions here if needed
		})
	}
}
