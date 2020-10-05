func maxSubArray(nums []int) int {
    numsLen := len(nums)
    if numsLen == 0 {
        return 0
    }
    if numsLen == 1 {
        return nums[0]
    }
    
    var maxEnd, maxSoFar = nums[0], nums[0]
    
    for i:=1; i<numsLen; i++ {
        if maxEnd < 0 && maxEnd < nums[i] {
            maxEnd = nums[i]
        } else {
            maxEnd += nums[i]
        }

        if maxEnd > maxSoFar {
            maxSoFar = maxEnd
        }
    }
    
    return maxSoFar
}