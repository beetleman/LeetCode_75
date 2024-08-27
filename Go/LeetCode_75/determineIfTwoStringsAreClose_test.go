package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseStrings(t *testing.T) {
	for _, scenario := range []struct {
		expected bool
		word1    string
		word2    string
	}{
		{
			expected: true,
			word1:    "abc",
			word2:    "bca",
		},
		{
			expected: false,
			word1:    "a",
			word2:    "aa",
		},
		{
			expected: true,
			word1:    "cabbba",
			word2:    "abbccc",
		},
		{
			expected: false,
			word1:    "uau",
			word2:    "ssx",
		},
	} {
		assert.Equal(t, scenario.expected, closeStrings(scenario.word1, scenario.word2))
	}
}
