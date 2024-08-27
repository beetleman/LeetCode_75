package leetcode75

import (
	"fmt"
	"slices"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	l1 := make(map[byte]int)
	l2 := make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		a := word1[i]
		b := word2[i]
		l1[a]++
		l2[b]++
	}

	if len(l1) != len(l2) {
		return false
	}

	vals1 := make([]int, 0, len(l1))
	vals2 := make([]int, 0, len(l1))

	fmt.Println(word1, l1, word2, l2)
	for key, val1 := range l1 {
		if val2, ok := l2[key]; ok {
			vals1 = append(vals1, val1)
			vals2 = append(vals2, val2)
		} else {
			return false
		}
	}

	fmt.Println(vals1, vals2)
	slices.Sort(vals1)
	slices.Sort(vals2)

	return slices.Equal(vals1, vals2)
}
