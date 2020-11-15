func numDecodings(s string) int {
	lenS := len(s)

	if lenS == 0 {
		return 0
	}

	var numPrevTwo, numPrevOne, num int
	numPrevTwo = 0

	if s[0] == '0' {
		numPrevTwo = 0
	} else {
		numPrevTwo = 1
	}

	if lenS == 1 {
		return numPrevTwo
	}

	if s[0] == '0' && s[1] == '0' {
		numPrevOne = 0
	} else if s[0] == '0' {
		numPrevOne = 1
	} else if s[0] == '1' || (s[0] == '2' && s[1] <= '6') {
		numPrevOne = 2
	} else {
		numPrevOne = 0
	}

	if lenS == 2 {
		return numPrevOne
	}

	num = 0

	for i := 2; i < lenS; i++ {
		if s[i-1] == '0' && s[i] == '0' {
			num = numPrevTwo
		} else if s[i-1] == '0' {
			num = numPrevOne
		} else if s[i-1] == '1' && s[i] == '0' {
			num = numPrevTwo
		} else if s[i-1] == '2' && s[i] == '0' {
			num = numPrevTwo
		}
		} else if s[i] != '0' && (s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6') {
			num = numPrevOne + numPrevTwo
		} else {
			num = numPrevTwo
		}

		numPrevTwo = numPrevOne
		numPrevOne = num
	}

	return num
}