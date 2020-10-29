func jump(nums []int) int {
	lenNums := len(nums)
	if lenNums <= 1 {
		return 0
	}

	// 已包含跳第一步
	// 这里包含一个隐含条件，即nums[0]>=1，不然如果为0，则这个人永远都跳不出去了。
	// 这也是为什么能直接使用 l = r 的条件
	var l, r, times int = 1, nums[0], 1

	for r < lenNums-1 {
		times += 1
		tmpMax := 0
		for i := l; i <= r; i++ {
			if tmpMax < i+nums[i] {
				tmpMax = i + nums[i]
			}
		}
		l = r
		r = tmpMax
	}

	return times
}