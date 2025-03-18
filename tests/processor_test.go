package tests

import (
	"go-reloaded/textprocessor"
	"testing"
)

func TestProcessText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"10 (bin) files", "2 files"},
		{"1E (hex) cars", "30 cars"},
		{"this is fun (up, 2)", "THIS IS fun"},
	}

	for _, tt := range tests {
		result := textprocessor.ProcessText(tt.input)
		if result != tt.expected {
			t.Errorf("got %q, expected %q", result, tt.expected)
		}
	}
}
