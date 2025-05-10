package adf2md

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ParseADF parses an ADF JSON string into a Node structure
func ParseADF(jsonStr string) (*Node, error) {
	if strings.TrimSpace(jsonStr) == "" {
		return nil, errors.New("empty JSON string")
	}

	var node Node
	err := json.Unmarshal([]byte(jsonStr), &node)
	if err != nil {
		return nil, fmt.Errorf("error parsing ADF JSON: %w", err)
	}

	// Validate that this is an ADF document
	if node.Type != "doc" {
		return nil, fmt.Errorf("invalid ADF format: root node must be 'doc', got '%s'", node.Type)
	}

	return &node, nil
}
