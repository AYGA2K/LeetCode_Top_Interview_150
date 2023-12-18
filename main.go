package main

import (
	"fmt"

	"github.com/AYGA2K/LeetCode_Top_Interview_150/leet"
)

func main() {
	arr := []int{2, 3, 1, 1, 4}
	arr2 := []int{3, 2, 1, 0, 4}
	fmt.Println(leet.CanJump(arr))
	fmt.Println(leet.CanJump(arr2))
}
