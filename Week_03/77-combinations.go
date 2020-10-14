// 涉及到拷贝，效率和空间都很差。需重做。
func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	var i, j int

	if k == 0 {
		return result
	}

	if k == 1 {
		for i = 1; i <= n; i++ {
			result = append(result, []int{i})
		}
		return result
	}

	for _, item := range combine(n, k-1) {
		for j = item[len(item)-1] + 1; j <= n; j++ {
			tmp := make([]int, 0)
			tmp = append(tmp, item...)
			tmp = append(tmp, j)
			result = append(result, tmp)
		}
	}

	return result
}