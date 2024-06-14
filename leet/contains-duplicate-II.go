package leet

import "math"

func ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[i]; !ok {
			m[v] = i
		} else {
			if math.Abs(float64(m[v]-i)) <= float64(k) {
				return true
			}
			m[v] = i
		}
	}
	return false
}
