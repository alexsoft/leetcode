package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder

	word1L := len(word1)
	word2L := len(word2)
	maxLen := max(word1L, word2L)

	for i := range maxLen {
		if i < word1L {
			sb.WriteByte(word1[i])
		}
		if i < word2L {
			sb.WriteByte(word2[i])
		}
	}

	return sb.String()
}
