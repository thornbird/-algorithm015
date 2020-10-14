import "sort"

func permuteUnique(nums []int) [][]int {

	res := make([][]int, 0)

	if len(nums) == 0 {
		return res
	}

	if len(nums) == 1 {
		res = append(res, nums)
	}

	sort.Ints(nums)

	for i, num := range nums {
		if i+1 < len(nums) && nums[i+1] == nums[i] {
			continue
		}
		tmpNums := make([]int, i)
		copy(tmpNums, nums[0:i])
		tmpNums = append(tmpNums, nums[i+1:]...)
		permuteRecords := permuteUnique(tmpNums)
		for _, permuteRec := range permuteRecords {
			res = append(res, append(permuteRec, num))
		}
	}
	return res
}