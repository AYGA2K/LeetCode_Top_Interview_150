package leet

import "fmt"

func hammingWeight(n int) int {
	weight := 0
	if n == 0 {
		return 0
	}

	var binaryStr string
	for n > 0 {
		remainder := n % 2
		if remainder == 1 {
			weight++
		}
		binaryStr = fmt.Sprintf("%d", remainder) + binaryStr
		n = n / 2
	}

	return weight
}
