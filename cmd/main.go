package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"text/template"

	templatesfs "github.com/zeek-r/trouble-tome/internal/fs"
)

type RunbookStep struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Runbook struct {
	Title string        `json:"title"`
	Steps []RunbookStep `json:"steps"`
}

const BinaryName = "troubletome"
const MarkDownTemplate = "markdown.tmpl"
const HTMLTemplate = "html.tmpl"
const DefaultOutputFile = "runbook"

var TemplatesDir = "templates"

var extensions = map[string]string{
	"markdown": "md",
	"html":     "html",
}

func main() {
	// Define command-line flags
	jsonFile := flag.String("json", "", "Path to the JSON source file")
	outputFormat := flag.String("format", "markdown", "Output format (markdown or html)")
	outputFile := flag.String("output", "runbook", "Path to the output file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "troubletome: A tool for generating incident management runbooks\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", BinaryName)
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -json runbook.json -format markdown -output runbook.md\n", BinaryName)
	}

	flag.Parse()

	if *jsonFile == "" {
		flag.PrintDefaults()
		log.Fatalf("Error: -json input is required")
	}

	templates := templatesfs.NewTemplates()

	// Read and parse JSON file
	runbook, err := readRunbookFromJSON(*jsonFile)
	if err != nil {
		log.Fatalf("Error reading input %s: %v", *jsonFile, err)
	}

	// Generate runbook based on format
	var output string
	switch strings.ToLower(*outputFormat) {
	case "markdown":
		output, err = generateMarkdownRunbook(runbook, templates)
	case "html":
		output, err = generateHTMLRunbook(runbook, templates)
	default:
		log.Fatalf("Unsupported output format: %s", *outputFormat)
	}

	if err != nil {
		log.Fatalf("Error generating runbook: %v", err)
	}

	// Write output to file]
	outputF := fmt.Sprintf("%s.%s", "runbook", extensions[*outputFormat])
	if outputFile != nil && *outputFile != "" && *outputFile != DefaultOutputFile {
		outputF = *outputFile
	}

	err = os.WriteFile(outputF, []byte(output), 0644)
	if err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	fmt.Printf("Runbook successfully generated: %s\n", outputF)
}

func readRunbookFromJSON(filePath string) (Runbook, error) {
	var runbook Runbook
	data, err := os.ReadFile(filePath)
	if err != nil {
		return runbook, err
	}
	err = json.Unmarshal(data, &runbook)
	return runbook, err
}

func generateMarkdownRunbook(runbook Runbook, templates fs.FS) (string, error) {
	return generate(MarkDownTemplate, runbook, templates)
}

func generateHTMLRunbook(runbook Runbook, templates fs.FS) (string, error) {
	return generate(HTMLTemplate, runbook, templates)
}

func generate(name string, runbook Runbook, templates fs.FS) (string, error) {
	tmpl, err := NewTemplate(name).ParseFS(templates, fmt.Sprintf("%s/%s", TemplatesDir, name))
	if err != nil {
		return "", err
	}

	var output strings.Builder
	err = tmpl.Execute(&output, runbook)
	return output.String(), err
}

func NewTemplate(name string) *template.Template {
	return template.New(name).Funcs(template.FuncMap{
		"inc":   func(i int) int { return i + 1 },
		"dec":   func(i int) int { return i - 1 },
		"slug":  func(s string) string { return strings.ToLower(strings.Replace(s, " ", "-", -1)) },
		"last":  func(x int, a []RunbookStep) bool { return x == len(a)-1 },
		"first": func(x int) bool { return x == 0 },
	})
}
