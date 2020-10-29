// 是否能由另一个字符串组成某个串
// 跟异位词的思路类似

func canConstruct(ransomNote string, magazine string) bool {
	canConstruct := true

	charMap := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		charMap[magazine[i]] = 1
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