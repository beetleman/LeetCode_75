package leetcode75

func maxOperations(nums []int, k int) int {
	operations := 0
	end := len(nums) - 1
	for i := 0; i <= end; i++ {
		if nums[i] >= k {
			continue
		}
		for j := i + 1; j <= end; j++ {
			if nums[i]+nums[j] == k {
				operations++
				nums[j] = nums[end]
				end--
				break
			}
		}
	}
	return operations
}
