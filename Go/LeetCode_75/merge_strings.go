package leetcode75

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	runes1, runes2 := []rune(word1), []rune(word2)

	var builder strings.Builder

	for i := 0; i < len(runes1) || i < len(runes2); i++ {
		if i < len(runes1) {
			builder.WriteRune(runes1[i])
		}
		if i < len(runes2) {
			builder.WriteRune(runes2[i])
		}
	}

	return builder.String()
}
