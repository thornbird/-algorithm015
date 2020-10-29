// 深度优先搜索
// 搜索终止条件：数组要么遍历完没有下一轮，要么就达到条件返回
// 这里充分利用题设 - 改变只用判断一个字符变化（字符串总由四个不同字符组合成），并且长度总是为8，可以省不少代码

const InitMin int = -1

func minMutation(start string, end string, bank []string) int {
	min := InitMin

	solution(start, end, bank, 0, &min)

	return min
}

func solution(current, end string, bank []string, count int, min *int) {
	if current == end {
		if *min == InitMin {
			*min = count
		} else if *min > count {
			*min = count
		}
		return
	}

	var tmpString string

	for i := 0; i < len(bank); i++ {
		if bank[i] == "" {
			continue
		}
		if oneDiff(current, bank[i]) {
			tmpString = bank[i]
			bank[i] = ""
			solution(tmpString, end, bank, count+1, min)
			bank[i] = tmpString
		}
	}
	return
}

// 这里s和t长度都为8 题设
func oneDiff(s, t string) bool {
	diffCount := 0
	for i := 0; i < 8; i++ {
		if s[i] != t[i] {
			diffCount++
		}
	}

	if diffCount == 1 {
		return true
	}

	return false
}