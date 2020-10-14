// 本题内存分配是个问题

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	if len(nums) == 0 {
		return res
	}

	if len(nums) == 1 {
		res = append(res, nums)
	}

	for i, num := range nums {
		tmpNums := make([]int, i)
		copy(tmpNums, nums[0:i])
		tmpNums = append(tmpNums, nums[i+1:]...)
		permuteRecords := permute(tmpNums)
		for _, permuteRec := range permuteRecords {
			res = append(res, append(permuteRec, num))
		}
	}
	return res
}