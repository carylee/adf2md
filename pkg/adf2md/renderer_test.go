package adf2md_test

import (
	"testing"

	"github.com/cary/adf2md/pkg/adf2md"
)

func TestRenderer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty document",
			input:    `{"version":1,"type":"doc","content":[]}`,
			expected: "",
		},
		{
			name:     "Simple paragraph",
			input:    `{"version":1,"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Hello, world!"}]}]}`,
			expected: "Hello, world!\n\n",
		},
		{
			name:     "Text with formatting",
			input:    `{"version":1,"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Plain"},{"type":"text","text":"Bold","marks":[{"type":"strong"}]},{"type":"text","text":"Italic","marks":[{"type":"em"}]}]}]}`,
			expected: "Plain**Bold***Italic*\n\n",
		},
		{
			name:     "Heading",
			input:    `{"version":1,"type":"doc","content":[{"type":"heading","attrs":{"level":2},"content":[{"type":"text","text":"Heading 2"}]}]}`,
			expected: "## Heading 2\n\n",
		},
		{
			name:     "Bullet list",
			input:    `{"version":1,"type":"doc","content":[{"type":"bulletList","content":[{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Item 1"}]}]},{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"Item 2"}]}]}]}]}`,
			expected: "* Item 1\n* Item 2\n",
		},
		{
			name:     "Code block",
			input:    `{"version":1,"type":"doc","content":[{"type":"codeBlock","attrs":{"language":"go"},"content":[{"type":"text","text":"func main() {\n\tfmt.Println(\"Hello\")\n}"}]}]}`,
			expected: "```go\nfunc main() {\n\tfmt.Println(\"Hello\")\n}\n```\n\n",
		},
	}

	renderer := adf2md.NewRenderer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, err := adf2md.ParseADF(tt.input)
			if err != nil {
				t.Fatalf("Failed to parse ADF: %v", err)
			}

			result, err := renderer.RenderToMarkdown(node)
			if err != nil {
				t.Fatalf("RenderToMarkdown failed: %v", err)
			}

			if result != tt.expected {
				t.Errorf("\nExpected: %q\nGot:      %q", tt.expected, result)
			}
		})
	}
}