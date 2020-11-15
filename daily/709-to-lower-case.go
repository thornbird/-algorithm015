func toLowerCase(str string) string {
	result := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if str[i] < 'a' && str[i] >= 'A' {
			result = append(result, str[i]+'a'-'A')
		} else {
			result = append(result, str[i])
		}
	}

	return string(result)
}

//  另一种写法也可以
func toLowerCase(str string) string {
	result := make([]rune, 0)
	for _, c := range str {
		if c < 'a' && c >= 'A' {
			result = append(result, c+'a'-'A')
		} else {
			result = append(result, c)
		}
	}

	return string(result)
}

// 还有一种写法，但这里居然没有报错，改了rune中的遍历c
func toLowerCase(str string) string {
	result := make([]rune, 0)
	for c := range str {
		if c < 'a' && c >= 'A' {
			result = append(result, c+'a'-'A')
		} else {
			result = append(result, c)
		}
	}

	return string(result)
}

// 但却不能直接用 str[i] = str[i] + 'a' - 'A' 这样