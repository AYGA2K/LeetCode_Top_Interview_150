package leet


func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	var afterString string
	for _, v := range s {
		if (v >= 97 && v <= 122) || (v >= 48 && v <= 57) {
			afterString = afterString + string(v)
			continue
		}
		if v >= 65 && v <= 90 {
			afterString = afterString + string(v+32)
		}
	}
	if len(afterString) <= 0 {
		return true
	}
	left, right := 0, len(afterString)-1
	for left < right {
		if afterString[left] != afterString[right] {
			return false
		}
		left++
		right--
	}
	return true
}
