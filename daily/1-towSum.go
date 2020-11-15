func twoSum(nums []int, target int) []int {
	res := make([]int, 0, 2)

	if len(nums) <= 1 {
		return res
	}

	tmpMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if idx, ok := tmpMap[target-nums[i]]; ok {
			res = append(res, []int{idx, i}...)
			break
		} else {
			tmpMap[nums[i]] = i
		}
	}

	return res
}
