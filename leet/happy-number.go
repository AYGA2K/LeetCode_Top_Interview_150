package leet

func IsHappy(n int) bool {
	m := make(map[int]bool)
	for {
		sum := squareSum(n)
		if sum == 1 {
			return true
		}
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = true
		n = sum
	}
}

func squareSum(n int) int {
	sum := 0
	for n != 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
