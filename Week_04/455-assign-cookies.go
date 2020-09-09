import "sort"
func findContentChildren(g []int, s []int) int {
    res := 0
    i:=0
    j:=0

    sort.Ints(g)
    sort.Ints(s)

    for i<len(g)&&j<len(s){
            if g[i] <= s[j] {
                res++
                i++
                j++
            } else {
                j++
            }
        }
    return res
}