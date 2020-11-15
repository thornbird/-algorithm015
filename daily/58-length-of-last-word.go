// 非常值得多写几次，理清楚这里面的index变化
func lengthOfLastWord(s string) int {
	var i, j int
	for j = len(s) - 1; j >= 0 && s[j] == ' '; j-- {
	}
	for i = j; i >= 0; i-- {
		if s[i] == ' ' {
			return j - i
		}
	}

	if j < 0 {
		return 0
	}

	return j + 1
}