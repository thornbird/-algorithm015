// Runtime 0ms; Memory 3MB

func isAnagram(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)

	if lenS != lenT {
		return false
	}

	if lenS == 0 {
		return true
	}

	helpArray := [26]int{}

	for _, item := range s {
		helpArray[item-'a']++
	}

	for _, item := range t {
		if helpArray[item-'a'] <= 0 {
			return false
		}
		helpArray[item-'a']--
	}

	return true
}