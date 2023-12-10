package main

import (
	"fmt"

	"github.com/AYGA2K/LeetCode_Top_Interview_150/leet"
)

func main() {
	strArray := []string{"hiasss", "hia", "hiafad"}
	fmt.Println(leet.LongestCommonPrefix(strArray))
	strArray2 := []string{"ab", "a"}
	fmt.Println(leet.LongestCommonPrefix(strArray2))
}
