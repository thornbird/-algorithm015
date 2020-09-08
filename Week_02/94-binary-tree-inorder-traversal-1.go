 /**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode	
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack1 := make([]interface{}, 0)
	cur := root

	for true{
		if isLeave(cur) || cur == nil {
			if cur != nil {
				res = append(res, cur.Val)
			}
			// 比如一个只有左子树的二叉树 会有出栈取出的节点右子树为空的情况
			if len(stack1) > 0 {
				cur, _ = (stack1[len(stack1)-1]).(*TreeNode)
                // 注意取出栈顶之后，要将栈顶元素给清理
                stack1 = stack1[:len(stack1)-1]
				res = append(res, cur.Val)
				cur = cur.Right
			} else {
				break
			}
		} else if cur.Left != nil {
			// 如果有左子树，则将当前节点压栈，并将遍历节点设置为左儿子
            stack1 = append(stack1, cur)
			cur = cur.Left
		} else {
            res = append(res, cur.Val)
			cur = cur.Right
		}
	}

	return res
}

func isLeave(cur *TreeNode) bool{
	return cur != nil && cur.Left == nil && cur.Right == nil
}