package leet

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverse(x)
}

func reverse(n int) int {
	if n > 0 {
		return n%10 + reverse(n/10)*10
	}
	return 0
}
