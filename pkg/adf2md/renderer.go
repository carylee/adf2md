package adf2md

import (
	"fmt"
	"strconv"
	"strings"
)

// Renderer handles converting ADF nodes to Markdown
type Renderer struct {
	// Options for customizing the Markdown output
	options RenderOptions
}

// RenderOptions contains configuration for the Markdown rendering
type RenderOptions struct {
	// Number of spaces used for list item indentation
	ListIndent int
}

// NewRenderer creates a new Markdown renderer with default options
func NewRenderer() *Renderer {
	return &Renderer{
		options: RenderOptions{
			ListIndent: 2, // Default to 2 spaces for list indentation
		},
	}
}

// WithOptions returns a new Renderer with the specified options
func (r *Renderer) WithOptions(options RenderOptions) *Renderer {
	r.options = options
	return r
}

// RenderToMarkdown converts an ADF node to Markdown
func (r *Renderer) RenderToMarkdown(node *Node) (string, error) {
	if node == nil {
		return "", fmt.Errorf("nil node provided")
	}
	
	result := r.renderNode(node)
	return result, nil
}

// renderNode processes a single ADF node and returns its Markdown representation
func (r *Renderer) renderNode(node *Node) string {
	if node == nil {
		return ""
	}

	switch node.Type {
	case "doc":
		return r.renderContent(node.Content)
	case "paragraph":
		return r.renderParagraph(node)
	case "text":
		return r.renderText(node)
	case "heading":
		return r.renderHeading(node)
	case "bulletList":
		return r.renderBulletList(node) + "\n"
	case "orderedList":
		return r.renderOrderedList(node) + "\n"
	case "listItem":
		return r.renderListItem(node)
	case "taskList":
		return r.renderTaskList(node) + "\n"
	case "taskItem":
		return r.renderTaskItem(node)
	case "decisionList":
		return r.renderDecisionList(node) + "\n"
	case "decisionItem":
		return r.renderDecisionItem(node)
	case "codeBlock":
		return r.renderCodeBlock(node)
	case "rule":
		return r.renderRule()
	case "blockquote":
		return r.renderBlockquote(node)
	case "hardBreak":
		return r.renderHardBreak()
	case "panel":
		return r.renderPanel(node)
	case "mention":
		return r.renderMention(node)
	case "emoji":
		return r.renderEmoji(node)
	case "date":
		return r.renderDate(node)
	case "status":
		return r.renderStatus(node)
	case "mediaSingle":
		return r.renderMediaSingle(node)
	case "media":
		return r.renderMedia(node)
	case "caption":
		return r.renderCaption(node)
	default:
		return r.renderUnknown(node)
	}
}

// renderContent processes an array of ADF nodes
func (r *Renderer) renderContent(nodes []Node) string {
	if len(nodes) == 0 {
		return ""
	}

	var result strings.Builder
	for _, node := range nodes {
		result.WriteString(r.renderNode(&node))
	}
	return result.String()
}

// renderParagraph renders a paragraph node
func (r *Renderer) renderParagraph(node *Node) string {
	content := r.renderContent(node.Content)
	if content == "" {
		return ""
	}
	return content + "\n\n"
}

// renderText renders a text node with any marks applied
func (r *Renderer) renderText(node *Node) string {
	if node.Text == "" {
		return ""
	}
	
	text := node.Text
	
	// Apply marks in reverse order (from most nested to least nested)
	if len(node.Marks) > 0 {
		for i := len(node.Marks) - 1; i >= 0; i-- {
			mark := node.Marks[i]
			switch mark.Type {
			case "strong":
				text = "**" + text + "**"
			case "em":
				text = "*" + text + "*"
			case "code":
				text = "`" + text + "`"
			case "strike":
				text = "~~" + text + "~~"
			case "underline":
				// Markdown doesn't natively support underline, could use HTML or just let it pass
				text = text // Could use "_" + text + "_" but that's italics in most Markdown
			case "link":
				if href, ok := mark.Attrs["href"].(string); ok {
					text = "[" + text + "](" + href + ")"
				}
			case "textColor":
				// Markdown doesn't support text color, could use HTML
				text = text
			case "backgroundColor":
				// Markdown doesn't support background color, could use HTML
				text = text
			}
		}
	}
	
	return text
}

// renderHeading renders a heading node
func (r *Renderer) renderHeading(node *Node) string {
	level := 1
	if lvl, ok := node.Attrs["level"].(float64); ok {
		level = int(lvl)
	}
	
	// Make sure level is between 1-6
	if level < 1 {
		level = 1
	} else if level > 6 {
		level = 6
	}
	
	content := r.renderContent(node.Content)
	return strings.Repeat("#", level) + " " + content + "\n\n"
}

// renderBulletList renders a bullet list node
func (r *Renderer) renderBulletList(node *Node) string {
	if len(node.Content) == 0 {
		return ""
	}
	
	var result strings.Builder
	for i, item := range node.Content {
		if i > 0 {
			result.WriteString("\n")
		}
		result.WriteString("* " + r.renderListItem(&item))
	}
	
	return result.String()
}

// renderOrderedList renders an ordered list node
func (r *Renderer) renderOrderedList(node *Node) string {
	if len(node.Content) == 0 {
		return ""
	}
	
	var result strings.Builder
	startOrder := 1
	if order, ok := node.Attrs["order"].(float64); ok {
		startOrder = int(order)
	}
	
	for i, item := range node.Content {
		if i > 0 {
			result.WriteString("\n")
		}
		num := startOrder + i
		result.WriteString(strconv.Itoa(num) + ". " + r.renderListItem(&item))
	}
	
	return result.String()
}

// renderListItem renders a list item node
func (r *Renderer) renderListItem(node *Node) string {
	if len(node.Content) == 0 {
		return "\n"
	}

	// Process the first child (usually a paragraph)
	firstChild := r.renderNode(&node.Content[0])
	firstChild = strings.TrimSuffix(firstChild, "\n\n") // Remove paragraph spacing
	
	// Process any additional content (nested lists, paragraphs, etc.)
	var result strings.Builder
	result.WriteString(firstChild)
	
	if len(node.Content) > 1 {
		indent := strings.Repeat(" ", r.options.ListIndent)
		
		for _, child := range node.Content[1:] {
			childContent := r.renderNode(&child)
			
			// For nested lists, we keep their formatting
			if child.Type == "bulletList" || child.Type == "orderedList" || child.Type == "taskList" {
				// Indent nested list items
				lines := strings.Split(childContent, "\n")
				for i, line := range lines {
					if line != "" {
						lines[i] = indent + line
					}
				}
				result.WriteString("\n" + strings.Join(lines, "\n"))
			} else {
				// For other block elements (paragraphs, etc.), indent them
				childContent = strings.TrimSuffix(childContent, "\n\n")
				lines := strings.Split(childContent, "\n")
				for i, line := range lines {
					if line != "" {
						lines[i] = indent + line
					}
				}
				result.WriteString("\n" + strings.Join(lines, "\n"))
			}
		}
	}
	
	return result.String()
}

// renderTaskList renders a task list node
func (r *Renderer) renderTaskList(node *Node) string {
	if len(node.Content) == 0 {
		return ""
	}
	
	var result strings.Builder
	for i, item := range node.Content {
		if i > 0 {
			result.WriteString("\n")
		}
		result.WriteString("- " + r.renderTaskItem(&item))
	}
	
	return result.String()
}

// renderTaskItem renders a task item node
func (r *Renderer) renderTaskItem(node *Node) string {
	state, _ := node.Attrs["state"].(string)
	checkbox := "[ ]"
	if state == "DONE" {
		checkbox = "[x]"
	}
	
	content := r.renderContent(node.Content)
	content = strings.TrimSuffix(content, "\n\n") // Remove paragraph spacing
	
	// Indent subsequent lines
	lines := strings.Split(content, "\n")
	for i := 1; i < len(lines); i++ {
		if lines[i] != "" {
			lines[i] = strings.Repeat(" ", r.options.ListIndent) + lines[i]
		}
	}
	
	return checkbox + " " + strings.Join(lines, "\n")
}

// renderDecisionList renders a decision list node
func (r *Renderer) renderDecisionList(node *Node) string {
	if len(node.Content) == 0 {
		return ""
	}
	
	var result strings.Builder
	for i, item := range node.Content {
		if i > 0 {
			result.WriteString("\n")
		}
		result.WriteString("- " + r.renderDecisionItem(&item))
	}
	
	return result.String()
}

// renderDecisionItem renders a decision item node
func (r *Renderer) renderDecisionItem(node *Node) string {
	state, _ := node.Attrs["state"].(string)
	prefix := "<D> "
	if state != "DECIDED" {
		prefix = "< > "
	}
	
	content := r.renderContent(node.Content)
	content = strings.TrimSuffix(content, "\n\n") // Remove paragraph spacing
	
	// Indent subsequent lines
	lines := strings.Split(content, "\n")
	for i := 1; i < len(lines); i++ {
		if lines[i] != "" {
			lines[i] = strings.Repeat(" ", r.options.ListIndent) + lines[i]
		}
	}
	
	return prefix + strings.Join(lines, "\n")
}

// renderCodeBlock renders a code block node
func (r *Renderer) renderCodeBlock(node *Node) string {
	language := ""
	if lang, ok := node.Attrs["language"].(string); ok {
		language = lang
	}
	
	var code string
	if len(node.Content) > 0 {
		code = node.Content[0].Text
	}
	
	return "```" + language + "\n" + code + "\n```\n\n"
}

// renderRule renders a horizontal rule
func (r *Renderer) renderRule() string {
	return "---\n\n"
}

// renderBlockquote renders a blockquote node
func (r *Renderer) renderBlockquote(node *Node) string {
	content := r.renderContent(node.Content)
	
	// Add blockquote prefix to each line
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = "> " + line
		}
	}
	
	return strings.Join(lines, "\n") + "\n\n"
}

// renderHardBreak renders a hard break (line break)
func (r *Renderer) renderHardBreak() string {
	return "  \n"
}

// renderPanel renders a panel node
func (r *Renderer) renderPanel(node *Node) string {
	panelType, _ := node.Attrs["panelType"].(string)
	title := fmt.Sprintf("**Panel (%s)**", panelType)
	
	content := r.renderContent(node.Content)
	if content != "" {
		content = title + "\n" + content
	} else {
		content = title
	}
	
	// Add blockquote prefix to each line
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = "> " + line
		}
	}
	
	return strings.Join(lines, "\n") + "\n\n"
}

// renderMention renders a mention node
func (r *Renderer) renderMention(node *Node) string {
	text, _ := node.Attrs["text"].(string)
	id, _ := node.Attrs["id"].(string)
	
	if text != "" {
		return "@" + text
	}
	return "@user:" + id
}

// renderEmoji renders an emoji node
func (r *Renderer) renderEmoji(node *Node) string {
	if text, ok := node.Attrs["text"].(string); ok {
		return text
	}
	if shortName, ok := node.Attrs["shortName"].(string); ok {
		return shortName
	}
	return ""
}

// renderDate renders a date node
func (r *Renderer) renderDate(node *Node) string {
	if timestamp, ok := node.Attrs["timestamp"].(string); ok {
		return "[Date: " + timestamp + "]"
	}
	return "[Date]"
}

// renderStatus renders a status node
func (r *Renderer) renderStatus(node *Node) string {
	if text, ok := node.Attrs["text"].(string); ok {
		return "[" + text + "]"
	}
	return "[STATUS]"
}

// renderMediaSingle renders a mediaSingle node
func (r *Renderer) renderMediaSingle(node *Node) string {
	if len(node.Content) == 0 {
		return ""
	}
	
	var result strings.Builder
	
	// First content item should be a media node
	if len(node.Content) > 0 && node.Content[0].Type == "media" {
		result.WriteString(r.renderMedia(&node.Content[0]))
	}
	
	// Second content item could be a caption
	if len(node.Content) > 1 && node.Content[1].Type == "caption" {
		caption := r.renderCaption(&node.Content[1])
		if caption != "" {
			result.WriteString("\n" + caption)
		}
	}
	
	return result.String()
}

// renderMedia renders a media node
func (r *Renderer) renderMedia(node *Node) string {
	mediaType, _ := node.Attrs["type"].(string)
	altText, _ := node.Attrs["alt"].(string)
	
	if altText == "" {
		altText = "image"
	}
	
	var url string
	if mediaType == "external" {
		url, _ = node.Attrs["url"].(string)
	} else {
		// For file or link types with collection/id
		id, _ := node.Attrs["id"].(string)
		collection, _ := node.Attrs["collection"].(string)
		url = "/wiki/download/attachments/" + collection + "/" + id
	}
	
	if url != "" {
		return "![" + altText + "](" + url + ")"
	}
	
	return "[Image: " + altText + " - Type: " + mediaType + "]"
}

// renderCaption renders a caption node
func (r *Renderer) renderCaption(node *Node) string {
	return "_" + r.renderContent(node.Content) + "_"
}

// renderUnknown handles unsupported node types
func (r *Renderer) renderUnknown(node *Node) string {
	return "[Unsupported ADF Element: " + node.Type + "]\n"
}

// indentSubsequentLines adds indentation to all lines after the first
func (r *Renderer) indentSubsequentLines(text string, indent string) string {
	lines := strings.Split(text, "\n")
	if len(lines) <= 1 {
		return text
	}
	
	for i := 1; i < len(lines); i++ {
		if lines[i] != "" {
			lines[i] = indent + lines[i]
		}
	}
	
	return strings.Join(lines, "\n")
}