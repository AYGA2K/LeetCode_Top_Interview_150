package leet

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if requiredIndx, ok := m[target-nums[i]]; ok {
			return []int{requiredIndx, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
