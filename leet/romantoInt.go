package leet

func romanToInt(s string) int {
	intNum := 0
	mapNum := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := 0; i < len(s); i++ {

		if i+1 < len(s) && mapNum[s[i]] < mapNum[s[i+1]] {
			intNum -= mapNum[s[i]]
			continue
		}
		if i+1 < len(s) && mapNum[s[i]] > mapNum[s[i+1]] {
			intNum += mapNum[s[i]]
			continue
		}
		intNum += mapNum[s[i]]
	}
	return intNum
}
