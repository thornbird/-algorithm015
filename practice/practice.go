package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0, 6)
	fmt.Printf("In main. Before passing: %p \n", &res)
	generate(0, 0, n, "", res)
	return res
}

func generate(l, r, n int, s string, res []string) {
	fmt.Printf("In general. Before passing: %p \n", &res)

	if l == n && r == n {
		res = append(res, s)
		fmt.Printf("In general. After passing: %p \n", &res)
		fmt.Println(res)
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

func main() {
	fmt.Println(generateParenthesis(3))
}
