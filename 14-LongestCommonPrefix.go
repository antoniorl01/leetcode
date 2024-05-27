package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var matches strings.Builder
	for letter_idx := 0; letter_idx < len(strs[0]); letter_idx++ {
		curr := strs[0][letter_idx]
		for word_idx := 1; word_idx < len(strs); word_idx++ {
			if letter_idx >= len(strs[word_idx]) || curr != strs[word_idx][letter_idx] {
				return matches.String()
			}
		}

		matches.WriteByte(curr)
	}

	return matches.String()
}
