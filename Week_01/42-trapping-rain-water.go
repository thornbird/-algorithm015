func trap(height []int) int {
    left := 0
    right := len(height) - 1
    maxLeft := 0
    maxRight := 0
    res := 0
    
    for left <= right {
        if height[left] <= height[right] {
            if height[left] >= maxLeft {
                maxLeft = height[left]
            } else {
                res += maxLeft - height[left]
            }
            left++
        } else {
            if height[right] >= maxRight {
                maxRight = height[right]
            } else {
                res += maxRight - height[right]
            }
            right--
        }
    }
    return res
}