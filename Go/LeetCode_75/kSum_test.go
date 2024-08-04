package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxOperations(t *testing.T) {
	for _, scenario := range []struct {
		nums   []int
		k      int
		output int
	}{
		{
			nums:   []int{1, 2, 3, 4},
			k:      5,
			output: 2,
		},
		{
			nums:   []int{3, 1, 3, 4, 3},
			k:      6,
			output: 1,
		},
	} {
		assert.Equal(t, scenario.output, maxOperations(scenario.nums, scenario.k))
	}
}
