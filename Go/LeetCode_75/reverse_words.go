package leetcode75

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	splited := strings.Fields(s)
	slices.Reverse(splited)
	return strings.Join(splited, " ")
}
