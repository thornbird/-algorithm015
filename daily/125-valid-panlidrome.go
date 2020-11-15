import (
	"strings"
)

// 注意 go里面没有while，用for实现
// for i < lenS && !numOrLetter(s[i]) index的判断放在前面
// 否则执行时按照&&的顺序，会出问题

func isPalindrome(s string) bool {
	lenS := len(s)

	if lenS <= 1 {
		return true
	}

	s = strings.ToUpper(s)

	i, j := 0, lenS-1

	for i < lenS && j >= 0 && i < j {
		for i < lenS && !numOrLetter(s[i]) {
			i++
		}

		for j >= 0 && !numOrLetter(s[j]) {
			j--
		}

		if i < lenS && j >= 0 {
			if s[i] == s[j] {
				i++
				j--
			} else {
				return false
			}
		}
	}

	return true
}

func numOrLetter(input byte) bool {
	if input >= '0' && input <= '9' {
		return true
	}

	if input >= 'A' && input <= 'Z' {
		return true
	}

	return false
}
