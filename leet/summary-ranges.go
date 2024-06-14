package leet

import "fmt"

func SummaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}
	start := nums[0]
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if start == end {
				result = append(result, fmt.Sprintf("%d", start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, end))
			}
			start = nums[i]
		}
		end = nums[i]
	}
	if start == end {
		result = append(result, fmt.Sprintf("%d", start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, end))
	}
	return result
}
