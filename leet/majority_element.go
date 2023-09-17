package leet

func MajorityElement(nums []int) int {
	major := 0
	index := 0
	count := 0
	for i, v := range nums {
		for _, val := range nums {
			if val == v {
				count++
			}
		}
		if count > major {
			major = count
			index = i
		}

		count = 0
	}
	return nums[index]
}
