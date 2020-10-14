package main

import "fmt"

func maxProfit(prices []int) int {
	lenPrices := len(prices)

	// 若股票价格天数小于等于1，返回0
	if lenPrices <= 1 {
		return 0
	}

	// 创建相邻差异数组
	diff := make([]int, 0, lenPrices-1)

	for i := 1; i < lenPrices; i++ {
		diff = append(diff, prices[i]-prices[i-1])
	}

	tmpProfit, maxProfit := 0, 0

	// 连续子序列的最大和
	for i := 0; i < lenPrices-1; i++ {
		if tmpProfit <= 0 {
			tmpProfit = diff[i]
		} else {
			tmpProfit += diff[i]
		}

		if tmpProfit > maxProfit {
			maxProfit = tmpProfit
		}
	}

	// 注意这里如果所有diff都<=0; 当不交易处理
	// 如果所有diff都<=0，还交易的话，就应该在初始赋值时选diff[0]
	return maxProfit
}

func main() {
	fmt.Println("Should equal 0", maxProfit([]int{}))
	fmt.Println("Should equal 0", maxProfit([]int{4}))
	fmt.Println("Should equal 0", maxProfit([]int{4, 4}))
	fmt.Println("Should equal 0", maxProfit([]int{4, 3, 2, 1}))

	fmt.Println("Should equal 3", maxProfit([]int{3, 1, 4, 2}))
	fmt.Println("Should equal 1", maxProfit([]int{3, 4, 1, 2}))
	fmt.Println("Should equal 6", maxProfit([]int{3, 1, 4, 6, 1, 7, 5}))
}
