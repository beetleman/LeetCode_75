package leetcode75

import (
	"math/bits"
)

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func minFlips(a int, b int, c int) int {
	maxLen := uint(max(
		bits.Len(uint(a)),
		bits.Len(uint(b)),
		bits.Len(uint(c)),
	))
	flips := 0

	for i := uint(0); i < maxLen; i++ {
		if hasBit(a, i) && hasBit(b, i) && !hasBit(c, i) {
			flips += 2
		} else if !hasBit(a, i) && !hasBit(b, i) && hasBit(c, i) {
			flips++
		} else if hasBit(a, i) && !hasBit(b, i) && !hasBit(c, i) {
			flips++
		} else if !hasBit(a, i) && hasBit(b, i) && !hasBit(c, i) {
			flips++
		}
	}

	return flips
}
