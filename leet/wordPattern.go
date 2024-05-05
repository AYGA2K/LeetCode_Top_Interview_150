package leet

import (
	"strings"
)

func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	m := make(map[string]string)
	m2 := make(map[string]string)

	for i, v := range pattern {
		if _, ok := m[string(v)]; !ok {
			m[string(v)] = words[i]
		} else if m[string(v)] != words[i] {
			return false
		}
	}
	for i, v := range words {
		if _, ok := m2[v]; !ok {
			m2[v] = string(pattern[i])
		}
	}
	return len(m) == len(m2)
}
