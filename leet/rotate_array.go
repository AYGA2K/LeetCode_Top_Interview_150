package leet

func Rotate(nums []int, k int) {
	k = len(nums) - k%len(nums)
	copy(nums, append(nums[k:], nums[:k]...))
}
