func numDistinct(s string, t string) int {
	lenS := len(s)
	lenT := len(t)

	if lenT == 0 {
		return 1
	}

	if lenT > lenS {
		return 0
	}

	helpDP := makeHelpDP(lenS, lenT)

	for j := 0; j <= lenS; j++ {
		helpDP[0][j] = 1
	}

	for i := 1; i <= lenT; i++ {
		for j := i; j <= lenS; j++ {
			if s[j-1] == t[i-1] {
				helpDP[i][j] = helpDP[i-1][j-1] + helpDP[i][j-1]
			} else {
				helpDP[i][j] = helpDP[i][j-1]
			}
		}
	}

	return helpDP[lenT][lenS]
}

func makeHelpDP(lenS, lenT int) [][]int {
	helpDP := make([][]int, lenT+1)
	for i := 0; i <= lenT; i++ {
		helpDP[i] = make([]int, lenS+1)
	}
	return helpDP
}