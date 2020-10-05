import "sort"

func fourSum(nums []int, target int) [][]int {
    res := make([][]int, 0)
    
    sort.Ints(nums)
    lenNums := len(nums)
    
    for i:=0; i<=lenNums-4; i++ {
        if i==0 || nums[i-1] != nums[i] {
            threeSums := threeSum(nums[i+1:lenNums], target-nums[i])
                for _, threeSum := range threeSums {
                    fourSum := append([]int{nums[i]}, threeSum...)
                    res = append(res, fourSum)
                }
        }        
    }
    
    return res
}

func threeSum(nums []int, target int) [][]int {
    res := make([][]int, 0)
    
    sort.Ints(nums)
    lenNums := len(nums)
    
    for i:=0; i<=lenNums-3; i++ {
        t := target - nums[i]
        l := i + 1
        r := lenNums - 1
        
        if i == 0 || nums[i-1] != nums[i] {
            for l < r {
                s := nums[l] + nums[r]
                if s == t {
                    res = append(res, []int{nums[i], nums[l], nums[r]})
                    for l < r && nums[l] == nums[l+1] { l+=1 }
                    for l < r && nums[r] == nums[r-1] { r-=1 }
                    l+=1
                    r-=1                    
                } else if s < t {
                    l+=1
                } else {
                    r-=1
                }
            }
        }
    }
    return res
}