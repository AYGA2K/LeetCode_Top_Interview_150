package leet

import "fmt"

func reverseBits(num uint32) uint32 {
	result := ""
	binary := fmt.Sprintf("%032b", num)

	for i := len(binary) - 1; i >= 0; i-- {
		result += string(binary[i])
	}
	var res uint32
	fmt.Sscanf(result, "%b", &res)
	return res
}
