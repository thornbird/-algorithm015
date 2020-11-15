/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 本题和 515 一样，都是用BFS，每次访问树的一层，再将下一层放进去
// 注意和模板的区别，每次一层，而非每次一个节点。
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	result := make([][]int, 0)
	var node *TreeNode

	for len(q) > 0 {
		curLen := len(q)
		curResult := make([]int, 0)

		for i := 0; i < curLen; i++ {
			node, q = q[0], q[1:]

			curResult = append(curResult, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		result = append(result, curResult)
	}

	return result
}
