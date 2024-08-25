package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOnes(t *testing.T) {
	for _, scenario := range []struct {
		expected int
		input    []int
		k        int
	}{
		{
			expected: 6,
			input:    []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
		},
		{
			expected: 10,
			input:    []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
		},
		{
			expected: 3,
			input:    []int{1, 1, 1},
			k:        3,
		},
	} {

		assert.Equal(t, scenario.expected, longestOnes(scenario.input, scenario.k))
	}

}
