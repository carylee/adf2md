package adf2md_test

import (
	"testing"

	"github.com/carylee/adf2md/pkg/adf2md"
)

func TestParseADF(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Empty input",
			input:   "",
			wantErr: true,
		},
		{
			name:    "Invalid JSON",
			input:   "{not json}",
			wantErr: true,
		},
		{
			name:    "Not ADF format",
			input:   `{"type":"notdoc"}`,
			wantErr: true,
		},
		{
			name:    "Valid minimal ADF",
			input:   `{"version":1,"type":"doc","content":[]}`,
			wantErr: false,
		},
		{
			name:    "Valid ADF with content",
			input:   `{"version":1,"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Hello"}]}]}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := adf2md.ParseADF(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseADF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
