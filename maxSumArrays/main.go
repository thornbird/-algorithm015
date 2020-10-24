package main

import "fmt"

func maxSumArrays(nums []int, len1, len2 int) int {
	lenNums := len(nums)
	if len1+len2 >= lenNums {
		return sumArray(nums)
	}

	maxResult := 0

	if len1 == len2 {
		diff := make([]int, lenNums-len1, lenNums-len1)
		for i := len1; i < lenNums; i++ {
			diff[i-len1] = nums[i] - nums[i-len1]
		}

		for i := len1; i <= lenNums-len2; i++ {
			maxSumArray1 := maxSumArray(nums[0:i], diff[0:i-len1], len1)
			maxSumArray2 := maxSumArray(nums[i:lenNums], diff[i:lenNums-len2], len2)
			if maxSumArray1+maxSumArray2 > maxResult {
				maxResult = maxSumArray1 + maxSumArray2
			}
		}

	} else {
		diff1 := make([]int, lenNums-len1, lenNums-len1)
		diff2 := make([]int, lenNums-len2, lenNums-len2)

		for i := len1; i < lenNums; i++ {
			diff1[i-len1] = nums[i] - nums[i-len1]
		}

		for i := len2; i < lenNums; i++ {
			diff2[i-len2] = nums[i] - nums[i-len2]
		}

		for i := len1; i <= lenNums-len2; i++ {
			maxSumArray1 := maxSumArray(nums[0:i], diff1[0:i-len1], len1)
			maxSumArray2 := maxSumArray(nums[i:lenNums], diff2[i:lenNums-len2], len2)
			if maxSumArray1+maxSumArray2 > maxResult {
				maxResult = maxSumArray1 + maxSumArray2
			}
		}

		for i := len2; i <= lenNums-len1; i++ {
			maxSumArray3 := maxSumArray(nums[0:i], diff2[0:i-len2], len2)
			maxSumArray4 := maxSumArray(nums[i:lenNums], diff1[i:lenNums-len1], len1)
			if maxSumArray3+maxSumArray4 > maxResult {
				maxResult = maxSumArray3 + maxSumArray4
			}
		}
	}
	return maxResult
}

// 算一个数组所有元素和
func sumArray(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

// 算一个数组中，连续sumLen个元素的最大和（这里已提供差值数组）
func maxSumArray(nums, diff []int, sumLen int) int {
	result := sumArray(nums[0:sumLen])

	cur, maxContinual := 0, 0
	for j := 0; j < len(diff); j++ {
		if cur >= 0 {
			cur += diff[j]
		} else {
			cur = diff[j]
		}
		if cur > maxContinual {
			maxContinual = cur
		}
	}

	return result + maxContinual
}

func main() {
	fmt.Println(maxSumArrays([]int{1, 4, 2, 5, 7, 3, 2}, 2, 2))
	fmt.Println(maxSumArrays([]int{1, 4, 2, 5, 7, 3, 2}, 4, 2))
}
