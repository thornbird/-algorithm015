func validPalindrome(s string) bool {
	s1 := []byte(s)

	i, j := 0, len(s1)-1

	for ; i < j && s1[i] == s1[j]; i++ {
		j--
	}

	if i >= j {
		return true
	}

	if isPalindrome(s1[i+1 : j+1]) {
		return true
	}

	if isPalindrome(s1[i:j]) {
		return true
	}

	return false
}

func isPalindrome(s []byte) bool {
	i, j := 0, len(s)-1

	for ; i < j && s[i] == s[j]; i++ {
		j--
	}

	if i >= j {
		return true
	}

	return false
}