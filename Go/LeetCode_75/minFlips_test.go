package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFlips(t *testing.T) {
	for _, scenario := range []struct {
		a      int
		b      int
		c      int
		output int
	}{
		{
			a:      2,
			b:      6,
			c:      5,
			output: 3,
		},
		{
			a:      2,
			b:      6,
			c:      5,
			output: 3,
		},
		{
			a:      1,
			b:      2,
			c:      3,
			output: 0,
		},
	} {
		assert.Equal(t, scenario.output, minFlips(scenario.a, scenario.b, scenario.c))
	}
}
