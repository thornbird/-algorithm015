package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0, 6)
	generate(0, 0, n, "", &res)
	return res
}

func generate(l, r, n int, s string, res *[]string) {
	if l == n && r == n {
		*res = append(*res, s)
		return
	}

	if l < n {
		fmt.Println(l)
		fmt.Println(r)
		generate(l+1, r, n, s+"(", res)
	}

	if r < l {
		fmt.Println(l)
		fmt.Println(r)
		generate(l, r+1, n, s+")", res)
	}

	return
}

func main() {
	fmt.Println(generateParenthesis(3))
}
