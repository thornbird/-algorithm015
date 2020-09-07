func twoSum(nums []int, target int) []int {
    diffMap := make(map[int]int)

    for index, value := range nums {        
        if diffIndex, exists := diffMap[target-value]; exists {
                return []int{diffIndex, index}
        }
        // 否则写入字典
        diffMap[value] = index
    }
    return []int{}

}