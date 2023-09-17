package leet

func RemoveDuplicates(nums []int) int {
	temp := nums[0]
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != temp {
			nums[length] = nums[i]
			length++
		}
		temp = nums[i]
	}
	return length
}
