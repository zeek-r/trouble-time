package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	templatesfs "github.com/zeek-r/trouble-tome/internal/fs"
)

func TestReadRunbookFromJSON(t *testing.T) {
	// Create a temporary JSON file
	content := `{
		"title": "Test Runbook",
		"steps": [
			{"title": "Step 1", "content": "Content 1"},
			{"title": "Step 2", "content": "Content 2"}
		]
	}`
	tmpfile, err := os.CreateTemp("", "test*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test readRunbookFromJSON
	runbook, err := readRunbookFromJSON(tmpfile.Name())
	if err != nil {
		t.Fatalf("readRunbookFromJSON failed: %v", err)
	}

	if runbook.Title != "Test Runbook" {
		t.Errorf("Expected title 'Test Runbook', got '%s'", runbook.Title)
	}
	if len(runbook.Steps) != 2 {
		t.Errorf("Expected 2 steps, got %d", len(runbook.Steps))
	}
}

func TestGenerateMarkdownRunbook(t *testing.T) {
	runbook := Runbook{
		Title: "Test Runbook",
		Steps: []RunbookStep{
			{Title: "Step 1", Content: "Content 1"},
			{Title: "Step 2", Content: "Content 2"},
		},
	}

	output, err := generateMarkdownRunbook(runbook, templatesfs.NewTemplates())
	if err != nil {
		t.Fatalf("generateMarkdownRunbook failed: %v", err)
	}

	expectedSubstrings := []string{
		"# Test Runbook",
		"## Step 1",
		"Content 1",
		"## Step 2",
		"Content 2",
	}

	for _, substr := range expectedSubstrings {
		if !strings.Contains(output, substr) {
			t.Errorf("Expected output to contain '%s', but it didn't", substr)
		}
	}
}

func TestGenerateHTMLRunbook(t *testing.T) {
	runbook := Runbook{
		Title: "Test Runbook",
		Steps: []RunbookStep{
			{Title: "Step 1", Content: "Content 1"},
			{Title: "Step 2", Content: "Content 2"},
		},
	}

	output, err := generateHTMLRunbook(runbook, templatesfs.NewTemplates())
	if err != nil {
		t.Fatalf("generateHTMLRunbook failed: %v", err)
	}

	expectedSubstrings := []string{
		"<title>Test Runbook</title>",
		"<h1>Test Runbook</h1>",
		"<h2 id=\"step-1\">Step 1</h2>",
		"Content 1",
		"<h2 id=\"step-2\">Step 2</h2>",
		"Content 2",
	}

	for _, substr := range expectedSubstrings {
		if !strings.Contains(output, substr) {
			t.Errorf("Expected output to contain '%s', but it didn't", substr)
		}
	}
}

func TestNewTemplate(t *testing.T) {
	// Create a test template string that uses all custom functions
	testTemplate := `
Title: {{.Title}}
Steps:
{{range $index, $step := .Steps}}
{{inc $index}}. {{$step.Title}}
   Content: {{$step.Content}}
   Index: {{$index}}
   Decremented Index: {{dec $index}}
   {{if first $index}}(First step){{end}}
   {{if last $index $.Steps}}(Last step){{end}}
Slug: {{slug $step.Title}}
{{end}}
`

	// Create a test runbook
	runbook := Runbook{
		Title: "Test Runbook",
		Steps: []RunbookStep{
			{Title: "Step One", Content: "Content 1"},
			{Title: "Step Two", Content: "Content 2"},
			{Title: "Final Step", Content: "Content 3"},
		},
	}

	// Create and parse the template
	tmpl := NewTemplate("test")
	parsedTmpl, err := tmpl.Parse(testTemplate)
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	// Execute the template
	var output strings.Builder
	err = parsedTmpl.Execute(&output, runbook)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	// Check the output
	expected := `
Title: Test Runbook
Steps:

1. Step One
   Content: Content 1
   Index: 0
   Decremented Index: -1
   (First step)
   Slug: step-one

2. Step Two
   Content: Content 2
   Index: 1
   Decremented Index: 0
   Slug: step-two
3. Final Step
   Content: Content 3
   Index: 2
   Decremented Index: 1

   (Last step)
   Slug: final-step`

	string1 := strings.ReplaceAll(output.String(), "\n", "")
	string2 := strings.ReplaceAll(expected, "\n", "")

	string1 = strings.ReplaceAll(string1, " ", "")
	string2 = strings.ReplaceAll(string2, " ", "")

	assert.Equal(t, string1, string2)
}
