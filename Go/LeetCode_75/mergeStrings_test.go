package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	for _, scenario := range []struct {
		expected string
		word1    string
		word2    string
	}{
		{
			expected: "apbqcr",
			word1:    "abc",
			word2:    "pqr",
		},
		{
			expected: "apbqrs",
			word1:    "ab",
			word2:    "pqrs",
		},
		{
			word1:    "abcd",
			word2:    "pq",
			expected: "apbqcd",
		},
	} {
		assert.Equal(t, scenario.expected, mergeAlternately(scenario.word1, scenario.word2))
	}
}
