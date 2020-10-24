// Runtime 8ms; Memory 3MB
func isAnagram(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)

	if lenS != lenT {
		return false
	}

	if lenS == 0 {
		return true
	}

	helpDict := make(map[byte]int)

	for i := 0; i < lenS; i++ {
		helpDict[s[i]]++
	}

	for i := 0; i < lenT; i++ {
		if times, ok := helpDict[t[i]]; ok {
			if times <= 0 {
				return false
			}
			helpDict[t[i]]--
		} else {
			return false
		}
	}
	return true
}