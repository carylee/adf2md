package adf2md

// ADF node types based on the Atlassian Document Format schema

// Node represents a basic ADF node structure
type Node struct {
	Type    string          `json:"type"`
	Version int             `json:"version,omitempty"`
	Content []Node          `json:"content,omitempty"`
	Text    string          `json:"text,omitempty"`
	Marks   []Mark          `json:"marks,omitempty"`
	Attrs   map[string]any  `json:"attrs,omitempty"`
}

// Mark represents formatting or other attributes applied to text
type Mark struct {
	Type  string         `json:"type"`
	Attrs map[string]any `json:"attrs,omitempty"`
}
