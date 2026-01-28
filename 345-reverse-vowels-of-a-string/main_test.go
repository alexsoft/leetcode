package main

import "testing"

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example 1",
			input:    "hello",
			expected: "holle",
		},
		{
			name:     "example 2",
			input:    "leetcode",
			expected: "leotcede",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no vowels",
			input:    "bcdfg",
			expected: "bcdfg",
		},
		{
			name:     "all vowels",
			input:    "aeiou",
			expected: "uoiea",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "uppercase vowels",
			input:    "hEllo",
			expected: "hollE",
		},
		{
			name:     "mixed case",
			input:    "aA",
			expected: "Aa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseVowels(tt.input)
			if result != tt.expected {
				t.Errorf("reverseVowels(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
