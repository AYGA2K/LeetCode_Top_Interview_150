package leet

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	for _, v := range s {
		map1[string(v)] = map1[string(v)] + 1
	}
	for _, v := range t {
		map2[string(v)] = map2[string(v)] + 1
	}
	for k, v := range map1 {
		if v != map2[k] {
			return false
		}
	}
	return true
}
