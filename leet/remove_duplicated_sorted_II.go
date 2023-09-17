package leet

func RemoveDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	length := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[length-2] {
			nums[length] = nums[i]
			length++
		}
	}
	return length
}
