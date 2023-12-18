package leet

func StrStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && len(needle) <= len(haystack)-i {
			for j := 0; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					break
				}
				if j == len(needle)-1 {
					return i
				}
			}
		}
	}
	return -1
}
