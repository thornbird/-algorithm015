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
				if visited {
					continue
				}
				for k := 0; k < len(q0); k++ {
					for l := 0; l < 26; l++ {
						if int(q0[k]-'a') == l {
							continue
						}
						tmp := []byte(q0)
						tmp[k] = byte(int('a') + l)
						if string(tmp) == wordList[j] {
							q = append(q, wordList[j])
						}
					}
				}

			}

		}
		step++
	}

	return 0
}