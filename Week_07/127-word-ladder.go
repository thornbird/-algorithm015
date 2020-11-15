func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == "" || endWord == "" || len(wordList) == 0 {
		return 0
	}

	q := []string{beginWord}
	step := 1
	used := make(map[string]bool)
	var q0 string

	for len(q) > 0 {
		lenQ := len(q)

		for i := 0; i < lenQ; i++ {
			q0, q = q[0], q[1:]

			if q0 == endWord {
				return step
			}

			for j := 0; j < len(wordList); j++ {
				_, visited := used[wordList[j]]
				if !visited && isOneDiff(q0, wordList[j]) {
					q = append(q, wordList[j])
					used[wordList[j]] = true
				}
			}
		}
		step++
	}

	return 0
}

func isOneDiff(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	j := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			j++
		}
		if j > 1 {
			return false
		}
	}

	if j == 1 {
		return true
	} else {
		return false
	}
}