package main

import (
	"fmt"

	"github.com/AYGA2K/LeetCode_Top_Interview_150/leet"
)

func main() {
	str := "   fly me   to   the moon  "
	str1 := "a "
	fmt.Println(leet.LengthOfLastWord(str))
	fmt.Println(leet.LengthOfLastWord(str1))
}
