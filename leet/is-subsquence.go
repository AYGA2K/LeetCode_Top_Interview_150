package leet

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	j := 0
	for _, v := range t {
		if byte(v) == s[j] {
			j++
		}
		if j == len(s) {
			return true
		}
	}
	return false
}
