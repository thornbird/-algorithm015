// 快排
func sortArray(nums []int) []int {
	arrLen := len(nums)
	if arrLen <= 1 {
		return nums
	}

	p := partition(nums, 0, arrLen-1)
	sortArray(nums[0:p])
	sortArray(nums[p+1 : arrLen])

	return nums
}

func partition(nums []int, low int, high int) (position int) {
	j := low
	for i := low + 1; i <= high; i++ {
		// 当前扫描值比哨兵小时，做交换
		if nums[i] < nums[low] && j+1 <= high {
			tmp := nums[i]
			j += 1
			nums[i] = nums[j]
			nums[j] = tmp
		}
	}

	if j != low {
		tmp := nums[j]
		nums[j] = nums[low]
		nums[low] = tmp
	}
	return j
}