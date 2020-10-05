func generateParenthesis(n int) []string {
    res := make([]string, 0)
    generate(0, 0, n, "", &res)
    return res
}

func generate(l, r, n int, s string, res *[]string){
    
    if l == n && r == n {
        *res = append(*res, s)
        return
    }
    
    if l < n {
        generate(l+1, r, n, s+"(", res)
    }
    
    if r < l {
        generate(l, r+1, n, s+")", res)
    }
    
    return
}
