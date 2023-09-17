package leet

import "fmt"

func Rotate(nums []int, k int) {
	k = len(nums) - k%len(nums)
	fmt.Println(k % len(nums))
	copy(nums, append(nums[k:], nums[:k]...))
	fmt.Println(nums)
}
