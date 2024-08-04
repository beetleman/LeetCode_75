package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	for _, scenario := range []struct {
		input  int
		output []int
	}{
		{
			input:  2,
			output: []int{0, 1, 1},
		},
		{
			input:  5,
			output: []int{0, 1, 1, 2, 1, 2},
		},
	} {
		assert.Equal(t, scenario.output, countBitsBuildins(scenario.input))
		assert.Equal(t, scenario.output, countBitsStr(scenario.input))

	}
}
