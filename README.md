# mdstrip

Strip Markdown formatting from text while preserving the content.

## Features

- Remove Markdown formatting syntax (headers, emphasis, lists, etc.)
- Process files or stdin
- Optional preservation of links and code blocks
- Fast and lightweight
- Cross-platform support

## Installation

### Using Homebrew

```bash
brew install daikw/tap/mdstrip
```

### Using Go

```bash
go install github.com/daikw/mdstrip/cmd@latest
```

### Download Binary

Download the latest release from the [releases page](https://github.com/daikw/mdstrip/releases).

## Usage

### Basic Usage

Strip Markdown from a file:
```bash
mdstrip input.md
```

Strip Markdown from stdin:
```bash
echo "# Hello **world**" | mdstrip
# Output: Hello world
```

Save output to a file:
```bash
mdstrip input.md -o output.txt
```

### Options

- `-o, --output` : Output file (default: stdout)
- `-l, --keep-links` : Keep link URLs in output
- `-c, --keep-code` : Keep code block markers
- `-V, --verbose` : Enable verbose logging
- `--version` : Show version information

### Examples

```bash
# Remove all Markdown formatting
echo "# Title\n**Bold** and *italic* text" | mdstrip
# Output: Title\nBold and italic text

# Keep links
echo "Visit [GitHub](https://github.com)" | mdstrip --keep-links
# Output: Visit GitHub (https://github.com)

# Process multiple files
for f in *.md; do mdstrip "$f" -o "${f%.md}.txt"; done
```

## What Gets Removed

- Headers (`# H1`, `## H2`, etc.)
- Emphasis (`**bold**`, `*italic*`, `__underline__`)
- Lists (`- item`, `* item`, `1. item`)
- Code blocks (` ```code``` `)
- Inline code (`` `code` ``)
- Links (`[text](url)`)
- Images (`![alt](url)`)
- Blockquotes (`> quote`)
- Horizontal rules (`---`, `***`)
- HTML tags
- Escape characters

## Development

### Prerequisites

- Go 1.22 or later
- Make (optional)

### Building

```bash
# Clone the repository
git clone https://github.com/daikw/mdstrip
cd mdstrip

# Download dependencies
make deps

# Build
make build

# Run tests
make test
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by various Markdown stripping tools in the ecosystem
- Built with [urfave/cli](https://github.com/urfave/cli) for CLI framework
- Uses [goreleaser](https://goreleaser.com/) for releases