package main

import (
	"slices"
	"testing"
)

func toSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func toList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "two elements",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := toList(tt.input)
			result := toSlice(reverseList(head))
			if !slices.Equal(result, tt.expected) {
				t.Errorf("reverseList(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
