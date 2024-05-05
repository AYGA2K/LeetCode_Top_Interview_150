package leet


func CanConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	count := 0
	length := len(ransomNote)
	usedLettersIndexes := []int{}
	for i := range ransomNote {
		for j := range magazine {
			if valueExists(usedLettersIndexes, i) {
				continue
			}
			if string(ransomNote[i]) == string(magazine[j]) {
				count++
				usedLettersIndexes = append(usedLettersIndexes, i)
				break
			}
		}
		if count == length {
			return true
		}
	}
	return false
}

func valueExists(array []int, value int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
