// 相比第二种做法，第一种做法申请了长和宽都少一的空间。
// 也更容易理解些，毕竟不是每个人都能想到空的字符串时，给长度赋值1
// 在赋值动态规划数组时，也考虑到当前s长度小于前t时，为0不用考虑这种情况，少了递推

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

	k := 0
	for j := 0; j < lenS; j++ {
		if s[j] == t[0] {
			k++
			helpDP[0][j] = k
		} else {
			helpDP[0][j] = helpDP[0][j-1]
		}
	}

	for i := 1; i < lenT; i++ {
		for j := i; j < lenS; j++ {
			if s[j] == t[i] {
				helpDP[i][j] = helpDP[i-1][j-1] + helpDP[i][j-1]
			} else {
				helpDP[i][j] = helpDP[i][j-1]
			}
		}
	}

	return helpDP[lenT-1][lenS-1]
}

func makeHelpDP(lenS, lenT int) [][]int {
	helpDP := make([][]int, lenT)
	for i := 0; i < lenT; i++ {
		helpDP[i] = make([]int, lenS)
	}
	return helpDP
}