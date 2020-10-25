// 广度优先搜索
// 在每次变化时，用到了bank

func minMutation(start string, end string, bank []string) int {
	changeMap := map[byte][]string{
		'A': []string{"C", "G", "T"},
		'C': []string{"A", "G", "T"},
		'G': []string{"C", "A", "T"},
		'T': []string{"C", "G", "A"},
	}

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

			for j := 0; j < 8; j++ {
				for k := 0; k < 3; k++ {
					// 每个字符对应另外三个可以换，参考changeMap
					tempMappingStr := changeMap[current[j]][k]
					// 经过一次变换（即替换一个字符）后的string
					tempStr := current[:j] + tempMappingStr + current[j+1:]
					if idx := idxOf(tempStr, bank); idx != -1 && !used[tempStr] {
						q = append(q, tempStr)
						used[tempStr] = true
					}
				}
			}
		}
		count++
	}

	return -1
}

func idxOf(s string, ss []string) int {
	for i := 0; i < len(ss); i++ {
		if s == ss[i] {
			return i
		}
	}
	return -1
}