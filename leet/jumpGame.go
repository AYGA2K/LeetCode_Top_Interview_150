package leet

func CanJump(nums []int) bool {
	maxReach := 0
	lastIndex := len(nums) - 1
	for i := 0; i <= maxReach && i <= lastIndex; i++ {
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= lastIndex {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
