func subsets(nums []int) [][]int {
	result := [][]int{[]int{}}

	for _, num := range nums {
		newSets := make([][]int, len(result))

		for i, item := range result {
			newSets[i] = make([]int, len(item))
			copy(newSets[i], item)
			newSets[i] = append(newSets[i], num)
		}

		result = append(result, newSets...)
	}

	return result
}