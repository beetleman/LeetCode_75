package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	for _, scenario := range []struct {
		s      string
		t      string
		output bool
	}{
		{
			s:      "abc",
			t:      "ahbgdc",
			output: true,
		},
		{
			s:      "axc",
			t:      "ahbgdc",
			output: false,
		},
		{
			s:      "",
			t:      "ahbgdc",
			output: true,
		},
	} {
		assert.Equal(t, scenario.output, isSubsequence(scenario.s, scenario.t))
	}
}
