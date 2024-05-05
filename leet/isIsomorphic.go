package leet

import "fmt"

func IsIsomorphic(s string, t string) bool {
	fmt.Println(s)
	fmt.Println(t)
	map1 := make(map[string]string)
	map2 := make(map[string]string)

	for i, v := range s {
		if _, ok := map1[string(v)]; !ok {
			map1[string(v)] = string(t[i])
		}else if map1[string(v)] != string(t[i]) {
			return false
		}
	}
	for i, v := range t {
		if _, ok := map2[string(v)]; !ok {
			map2[string(v)] = string(s[i])
		}
	}
	return len(map1) == len(map2)
}
