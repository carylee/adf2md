package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/carylee/adf2md/pkg/adf2md"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("file", "", "Input file containing ADF JSON (default: stdin)")
	outputFile := flag.String("output", "", "Output file for Markdown (default: stdout)")
	listIndent := flag.Int("indent", 2, "Number of spaces used for list item indentation")
	flag.Parse()

	// Get input content
	var input []byte
	var err error
	
	if *inputFile != "" {
		// Read from file
		input, err = os.ReadFile(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Check if there's extra argument as content
		if flag.NArg() > 0 {
			input = []byte(flag.Arg(0))
		} else {
			// Read from stdin
			input, err = readStdin()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
				os.Exit(1)
			}
		}
	}

	// Parse ADF JSON
	node, err := adf2md.ParseADF(string(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing ADF: %v\n", err)
		os.Exit(1)
	}

	// Convert to Markdown
	renderer := adf2md.NewRenderer().WithOptions(adf2md.RenderOptions{
		ListIndent: *listIndent,
	})
	
	markdown, err := renderer.RenderToMarkdown(node)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering Markdown: %v\n", err)
		os.Exit(1)
	}

	// Write output
	if *outputFile != "" {
		err = os.WriteFile(*outputFile, []byte(markdown), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(markdown)
	}
}

// readStdin reads all content from stdin
func readStdin() ([]byte, error) {
	// Check if stdin has data
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Read from stdin pipe
		reader := bufio.NewReader(os.Stdin)
		var builder strings.Builder
		
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			builder.WriteString(line)
		}
		
		return []byte(builder.String()), nil
	}
	
	return nil, fmt.Errorf("no input provided via stdin")
}
