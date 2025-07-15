package main

import (
	"testing"
)

func TestStripMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		opts     StripOptions
		expected string
	}{
		{
			name:     "Headers",
			input:    "# Header 1\n## Header 2\n### Header 3",
			opts:     StripOptions{},
			expected: "Header 1\nHeader 2\nHeader 3",
		},
		{
			name:     "Emphasis",
			input:    "**bold** and *italic* and __underline__ and _underscore_",
			opts:     StripOptions{},
			expected: "bold and italic and underline and underscore",
		},
		{
			name:     "Lists",
			input:    "- Item 1\n* Item 2\n+ Item 3\n1. Numbered 1\n2. Numbered 2",
			opts:     StripOptions{},
			expected: "Item 1\nItem 2\nItem 3\nNumbered 1\nNumbered 2",
		},
		{
			name:     "Links without keep",
			input:    "Visit [GitHub](https://github.com) for more",
			opts:     StripOptions{KeepLinks: false},
			expected: "Visit GitHub for more",
		},
		{
			name:     "Links with keep",
			input:    "Visit [GitHub](https://github.com) for more",
			opts:     StripOptions{KeepLinks: true},
			expected: "Visit GitHub (https://github.com) for more",
		},
		{
			name:     "Inline code",
			input:    "Use `fmt.Println()` to print",
			opts:     StripOptions{},
			expected: "Use fmt.Println() to print",
		},
		{
			name:     "Code blocks without keep",
			input:    "```go\nfunc main() {\n}\n```",
			opts:     StripOptions{KeepCode: false},
			expected: "func main() {\n}",
		},
		{
			name:     "Code blocks with keep",
			input:    "```go\nfunc main() {\n}\n```",
			opts:     StripOptions{KeepCode: true},
			expected: "```go\nfunc main() {\n}\n```",
		},
		{
			name:     "Blockquotes",
			input:    "> This is a quote\n> Another line",
			opts:     StripOptions{},
			expected: "This is a quote\nAnother line",
		},
		{
			name:     "Images",
			input:    "![Alt text](image.png) in the middle",
			opts:     StripOptions{},
			expected: "Alt text in the middle",
		},
		{
			name:     "Horizontal rules",
			input:    "Text above\n---\nText below",
			opts:     StripOptions{},
			expected: "Text above\n\nText below",
		},
		{
			name:     "HTML tags",
			input:    "Some <strong>HTML</strong> content",
			opts:     StripOptions{},
			expected: "Some HTML content",
		},
		{
			name:     "Escape characters",
			input:    "Escape \\* asterisk \\# hash",
			opts:     StripOptions{},
			expected: "Escape * asterisk # hash",
		},
		{
			name:     "Complex document",
			input:    "# Title\n\nThis is **bold** and *italic* text.\n\n- List item\n- Another item\n\n```\ncode block\n```\n\n> Quote\n\n[Link](url)",
			opts:     StripOptions{},
			expected: "Title\n\nThis is bold and italic text.\n\nList item\nAnother item\n\ncode block\n\nQuote\n\nLink",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StripMarkdown(tt.input, tt.opts)
			if result != tt.expected {
				t.Errorf("StripMarkdown() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestStripLineMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		opts     StripOptions
		expected string
	}{
		{
			name:     "Empty line",
			input:    "",
			opts:     StripOptions{},
			expected: "",
		},
		{
			name:     "Plain text",
			input:    "Just plain text",
			opts:     StripOptions{},
			expected: "Just plain text",
		},
		{
			name:     "Mixed formatting",
			input:    "# Header with **bold** and *italic*",
			opts:     StripOptions{},
			expected: "Header with bold and italic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stripLineMarkdown(tt.input, tt.opts)
			if result != tt.expected {
				t.Errorf("stripLineMarkdown() = %q, want %q", result, tt.expected)
			}
		})
	}
}