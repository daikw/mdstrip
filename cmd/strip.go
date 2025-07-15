package main

import (
	"regexp"
	"strings"
)

// StripOptions configures the markdown stripping behavior
type StripOptions struct {
	KeepLinks bool
	KeepCode  bool
}

// StripMarkdown removes markdown formatting from the input text
func StripMarkdown(input string, opts StripOptions) string {
	// Process line by line to handle block elements properly
	lines := strings.Split(input, "\n")
	var result []string
	inCodeBlock := false
	codeBlockMarker := ""

	for _, line := range lines {
		// Handle code blocks
		if strings.HasPrefix(strings.TrimSpace(line), "```") || strings.HasPrefix(strings.TrimSpace(line), "~~~") {
			if !inCodeBlock {
				inCodeBlock = true
				codeBlockMarker = strings.TrimSpace(line)[:3]
				if opts.KeepCode {
					result = append(result, line)
				}
				continue
			} else if strings.HasPrefix(strings.TrimSpace(line), codeBlockMarker) {
				inCodeBlock = false
				if opts.KeepCode {
					result = append(result, line)
				}
				continue
			}
		}

		// If in code block, keep content as is
		if inCodeBlock {
			if opts.KeepCode {
				result = append(result, line)
			} else {
				result = append(result, line)
			}
			continue
		}

		// Process the line
		processed := stripLineMarkdown(line, opts)
		if processed != "" || line == "" {
			result = append(result, processed)
		}
	}

	// Join and clean up
	output := strings.Join(result, "\n")
	
	// Clean up multiple blank lines
	output = regexp.MustCompile(`\n{3,}`).ReplaceAllString(output, "\n\n")
	
	return strings.TrimSpace(output)
}

// stripLineMarkdown removes markdown formatting from a single line
func stripLineMarkdown(line string, opts StripOptions) string {
	// Skip horizontal rules
	if regexp.MustCompile(`^\s*[-*_]{3,}\s*$`).MatchString(line) {
		return ""
	}

	// Remove headers (# ## ### etc)
	line = regexp.MustCompile(`^#{1,6}\s+`).ReplaceAllString(line, "")

	// Remove blockquotes
	line = regexp.MustCompile(`^>\s*`).ReplaceAllString(line, "")

	// Remove list markers (-, *, +, 1., 2., etc)
	line = regexp.MustCompile(`^[\s]*[-*+]\s+`).ReplaceAllString(line, "")
	line = regexp.MustCompile(`^[\s]*\d+\.\s+`).ReplaceAllString(line, "")

	// Remove emphasis (bold and italic)
	// Handle bold first (** or __)
	line = regexp.MustCompile(`\*\*([^*]+)\*\*`).ReplaceAllString(line, "$1")
	line = regexp.MustCompile(`__([^_]+)__`).ReplaceAllString(line, "$1")
	
	// Handle italic (* or _)
	line = regexp.MustCompile(`\*([^*]+)\*`).ReplaceAllString(line, "$1")
	line = regexp.MustCompile(`_([^_]+)_`).ReplaceAllString(line, "$1")

	// Remove strikethrough
	line = regexp.MustCompile(`~~([^~]+)~~`).ReplaceAllString(line, "$1")

	// Remove inline code
	line = regexp.MustCompile("`([^`]+)`").ReplaceAllString(line, "$1")

	// Handle links
	if opts.KeepLinks {
		// Keep both text and URL: [text](url) -> text (url)
		line = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`).ReplaceAllString(line, "$1 ($2)")
	} else {
		// Keep only text: [text](url) -> text
		line = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`).ReplaceAllString(line, "$1")
	}

	// Remove reference-style links
	line = regexp.MustCompile(`\[([^\]]+)\]\[[^\]]*\]`).ReplaceAllString(line, "$1")
	line = regexp.MustCompile(`^\[[^\]]+\]:\s*.*$`).ReplaceAllString(line, "")

	// Remove images
	line = regexp.MustCompile(`!\[([^\]]*)\]\([^)]+\)`).ReplaceAllString(line, "$1")

	// Remove HTML tags
	line = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(line, "")

	// Remove escape characters
	line = regexp.MustCompile(`\\([\\`*_{}\[\]()#+\-.!])`).ReplaceAllString(line, "$1")

	return strings.TrimSpace(line)
}