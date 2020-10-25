/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	helpQueue := []*TreeNode{root}
	result := make([]int, 0)
	var node *TreeNode

	for len(helpQueue) > 0 {
		curLevelLen := len(helpQueue)
		maxValue := math.MinInt64

		for i := 0; i < curLevelLen; i++ {
			node, helpQueue = helpQueue[0], helpQueue[1:]
			if node.Val > maxValue {
				maxValue = node.Val
			}

			if node.Left != nil {
				helpQueue = append(helpQueue, node.Left)
			}

			if node.Right != nil {
				helpQueue = append(helpQueue, node.Right)
			}
		}
		result = append(result, maxValue)
	}

	return result
}