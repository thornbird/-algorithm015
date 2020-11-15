func longestPalindrome(s string) string {
	lenS := len(s)
	if lenS < 2 {
		return s
	}
	// 中心向外扩张
	maxStart, maxLen := 0, 0

	for i := 0; i < lenS-1; i++ {
		extendPalindrome(s, i, i, &maxStart, &maxLen)
		extendPalindrome(s, i, i+1, &maxStart, &maxLen)
	}

	return string(s[maxStart : maxStart+maxLen])
}

func extendPalindrome(s string, start, end int, maxStart, maxLen *int) {
	for start >= 0 && end < len(s) {
		if s[start] == s[end] {
			start--
			end++
		} else {
			break
		}
	}

	if end-start-1 > *maxLen {
		*maxStart = start + 1
		*maxLen = end - start - 1
	}

	return
}