package leetcode75

func longestOnes(nums []int, k int) int {
	from, to := 0, 0
	zeros := 0
	for to < len(nums) {
		if nums[to] == 0 {
			zeros++
		}
		to++
		if k < zeros {
			if nums[from] == 0 {
				zeros--
			}
			from++
		}
	}
	return to - from
}
