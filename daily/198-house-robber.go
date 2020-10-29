// https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
	lenNums := len(nums)

	if lenNums == 0 {
		return 0
	}

	dp := make([]int, lenNums)
	dp[0] = nums[0]

	if lenNums > 1 {
		dp[1] = max(nums[0], nums[1])
	}

	for i := 2; i < lenNums; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[lenNums-1]
}

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}