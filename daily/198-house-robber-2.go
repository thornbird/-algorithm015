// 打家劫舍 用二维的dp数组来求解
// 注意这里的 dp[i][0] = max(dp[i-1][0], dp[i-1][1])
// 不能直接记录为 dp[i][0] = dp[i-1][1],  举例子 2112
// 可以连续两家不偷，不偷的两家最左和最右都偷
// 但不能连续三家不偷，题设中每家都有一定量的现金，三家连续不偷中这种场景一定小于偷三家中间的场景。

func rob(nums []int) int {
	lenNums := len(nums)

	if lenNums == 0 {
		return 0
	}

	if lenNums == 1 {
		return nums[0]
	}

	dp := make([][2]int, lenNums)

	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < lenNums; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return max(dp[lenNums-1][0], dp[lenNums-1][1])
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}