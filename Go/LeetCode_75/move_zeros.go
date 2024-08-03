package leetcode75

func moveZeroes(nums []int) {
	pivot := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pivot], nums[i] = nums[i], nums[pivot]
			pivot++
		}
	}
}
