func majorityElement(nums []int) int {
	var num, count = nums[0], 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			num, count = nums[i], 1
		} else if nums[i] == num {
			count++
		} else {
			count--
		}
	}

	return num
}