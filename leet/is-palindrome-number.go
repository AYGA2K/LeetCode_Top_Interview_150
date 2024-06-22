package leet

func IsPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverse(x)
}

func reverse(n int) int {
	reversed := 0
	for n > 0 {
		reversed = n%10 + reversed*10
		n /= 10
	}
	return reversed
}
