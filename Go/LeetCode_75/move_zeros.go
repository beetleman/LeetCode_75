package leetcode75

func moveZeroes(nums []int) {
	zeros := 0
	done := false

	for !done {
		done = true
		for i := 0; i+zeros < len(nums); i++ {
			if nums[i] == 0 {
				for j := i + 1; j+zeros < len(nums); j++ {
					nums[j-1] = nums[j]
				}
				nums[len(nums)-zeros-1] = 0
				zeros++
				done = false
			}
		}
	}
}
