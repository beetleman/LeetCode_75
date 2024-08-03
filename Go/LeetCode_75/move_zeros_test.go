package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	for _, scenario := range []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{0, 1, 0, 3, 12},
			output: []int{1, 3, 12, 0, 0},
		},
		{
			input:  []int{0},
			output: []int{0},
		},
		{
			input:  []int{0, 1, 0},
			output: []int{1, 0, 0},
		},
		{
			input:  []int{0, 0, 1},
			output: []int{1, 0, 0},
		},
	} {
		moveZeroes(scenario.input)
		assert.Equal(t, scenario.output, scenario.input)
	}
}
