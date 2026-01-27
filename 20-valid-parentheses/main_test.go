package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "simple parentheses",
			input:    "()",
			expected: true,
		},
		{
			name:     "multiple types",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "nested brackets",
			input:    "{[()]}",
			expected: true,
		},
		{
			name:     "mismatched brackets",
			input:    "(]",
			expected: false,
		},
		{
			name:     "unclosed bracket",
			input:    "([",
			expected: false,
		},
		{
			name:     "extra closing bracket",
			input:    "()]",
			expected: false,
		},
		{
			name:     "one closing bracket",
			input:    "]",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.input)
			if result != tt.expected {
				t.Errorf("isValid(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
