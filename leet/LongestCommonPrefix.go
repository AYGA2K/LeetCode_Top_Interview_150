package leet

func LongestCommonPrefix(strs []string) string {
	Common := ""
	firstWord := strs[0]
	for i := 0; i < len(firstWord); i++ {
		for j := 1; j < len(strs); j++ {
			if i > (len(strs[j]) - 1) {
				return Common
			}
			if i < len(strs[j]) && firstWord[i] != strs[j][i] {
				return Common
			}
		}
		Common = Common + string(firstWord[i])
	}

	return Common
}
