# adf2md

A Go command-line tool to convert Atlassian Document Format (ADF) JSON into Markdown.

## Overview

`adf2md` is a simple utility that converts Atlassian Document Format (ADF) JSON into standard Markdown. It handles most common ADF node types including text formatting, lists, tables, code blocks, and more.

## Features

- Convert ADF JSON to clean, readable Markdown
- Accept input from file, stdin, or command-line argument
- Output to file or stdout
- Simple, intuitive command-line interface

## Installation

### Homebrew

The easiest way to install `adf2md` is via Homebrew:

```bash
# Add the tap
brew tap carylee/adf2md

# Install the tool
brew install adf2md
```

### Go Install

If you have Go installed:

```bash
go install github.com/carylee/adf2md/cmd/adf2md@latest
```

### Binary Release

Alternatively, download a binary release from the [GitHub releases page](https://github.com/carylee/adf2md/releases).

## Usage

```bash
# Basic usage (reads from stdin, writes to stdout)
cat adf.json | adf2md

# Read from a file
adf2md -i input.json
# or with long form flag
adf2md --input input.json

# Write output to a file
adf2md -i input.json -o result.md
# or with long form flags
adf2md --input input.json --output result.md

# Pass ADF JSON directly as an argument
adf2md '{"version":1,"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Hello world"}]}]}'

# Get version information
adf2md -v
# or
adf2md --version

# Show help
adf2md -h
# or
adf2md --help
```

## Supported ADF Elements

- Document structure (`doc`)
- Paragraphs (`paragraph`)
- Headings (`heading` levels 1-6)
- Text (`text`) with marks:
  - Strong/Bold (`strong`)
  - Emphasis/Italic (`em`)
  - Inline Code (`code`)
  - Strike-through (`strike`)
  - Links (`link`)
- Lists:
  - Bullet Lists (`bulletList`)
  - Ordered Lists (`orderedList`)
  - Task Lists (`taskList`)
  - Decision Lists (`decisionList`)
- Code Blocks (`codeBlock`) with language specification
- Blockquotes (`blockquote`)
- Panels (`panel`) (rendered as styled blockquotes)
- Horizontal Rules (`rule`)
- Hard Breaks (`hardBreak`)
- Inline nodes:
  - Mentions (`mention`)
  - Emoji (`emoji`)
  - Date (`date`)
  - Status (`status`)
- Media: 
  - Media Single (`mediaSingle`)
  - Media (`media`) (basic image support)
  - Captions (`caption`)

## License

MIT

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.
