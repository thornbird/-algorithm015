func maxProfit(prices []int) int {
	lenPrices := len(prices)

	if lenPrices <= 1 {
		return 0
	}

	var maxProfit int

	for i := 1; i < lenPrices; i++ {
		if prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}