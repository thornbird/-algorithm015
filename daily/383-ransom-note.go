func canConstruct(ransomNote string, magazine string) bool {
	canConstruct := true

	charMap := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		if _, ok := charMap[magazine[i]]; ok {
			charMap[magazine[i]]++
		} else {
			charMap[magazine[i]] = 1
		}
	}

	for i := 0; i < len(ransomNote); i++ {
		if times, ok := charMap[ransomNote[i]]; ok {
			if times == 0 {
				canConstruct = false
				break
			} else {
				charMap[ransomNote[i]]--
			}
		} else {
			canConstruct = false
			break
		}
	}

	return canConstruct
}