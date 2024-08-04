package leetcode75

import (
	"strconv"
	"strings"

	"math/bits"
)

func countBitsBuildins(n int) []int {
	results := make([]int, n+1)
	for i := uint(0); i <= uint(n); i++ {
		results[i] = bits.OnesCount(i)
	}
	return results
}

func countBitsStr(n int) []int {
	results := make([]int, n+1)
	for i := 0; i <= n; i++ {
		results[i] = strings.Count(strconv.FormatInt(int64(i), 2), "1")
	}

	return results
}
