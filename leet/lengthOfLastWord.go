package leet

func LengthOfLastWord(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	end := length - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}

	lastWordLength := 0
	for i := end; i >= 0; i-- {
		if s[i] != ' ' {
			lastWordLength++
		} else {
			break
		}
	}

	return lastWordLength
}
