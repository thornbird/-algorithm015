// 广度优先搜索，充分利用题设。
// 直接看bank中的每个串和比较的串是不是相差一个字符。

func minMutation(start string, end string, bank []string) int {

	q := []string{start}
	used := make(map[string]bool)
	count := 0

	var current string

	for len(q) > 0 {
		lenQ := len(q)
		for i := 0; i < lenQ; i++ {
			current, q = q[0], q[1:]
			if current == end {
				return count
			}

			for i := 0; i < len(bank); i++ {
				if !used[bank[i]] && oneDiff(bank[i], current) {
					q = append(q, bank[i])
					used[bank[i]] = true
				}
			}
		}
		count++
	}

	return -1
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