package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWrolds(t *testing.T) {
	for _, scenario := range []struct {
		input  string
		output string
	}{
		{
			input:  "the sky is blue",
			output: "blue is sky the",
		},
		{
			input:  "  hello world  ",
			output: "world hello",
		},
		{
			input:  "a good   example",
			output: "example good a",
		},
	} {
		assert.Equal(t, scenario.output, reverseWords(scenario.input))
	}
}
