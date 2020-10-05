func jump(nums []int) int {
    lenNums := len(nums)
    if lenNums <= 1 {
        return 0
    }
    
    var l, r, times int = 1, nums[0], 1
    
    for r < lenNums - 1 {
        times += 1
        tmpMax := 0
        for i:=l; i<=r; i++ {
            if tmpMax < i + nums[i] {
                tmpMax = i + nums[i]
            }
        }
        l = r
        r = tmpMax
    }
    
    return times
}
