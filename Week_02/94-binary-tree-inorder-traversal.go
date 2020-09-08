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
		// 如果遍历到的节点为空
		if cur == nil {
			// 比如一个只有左子树的二叉树 会有出栈取出的节点右子树为空的情况
			if len(stack1) > 0 {
				cur, _ = (stack1[len(stack1)-1]).(*TreeNode)
                // 注意取出栈顶之后，要将栈顶元素给清理
                stack1 = stack1[:len(stack1)-1]
				res = append(res, cur.Val)
				cur = cur.Right
			} else {
				return res
			}
		} else if cur.Left != nil {
            // 如果有左子树，则将当前节点压栈，并将遍历节点设置为左儿子
            stack1 = append(stack1, cur)
			cur = cur.Left
		} else {
			// 打印当前节点
			res = append(res, cur.Val)
			// 有右子树，将遍历节点设置为右儿子
			if cur.Right != nil {
				cur = cur.Right
			} else {
				// 已是叶子节点，出栈
				if len(stack1) > 0 {
					cur, _ = (stack1[len(stack1)-1]).(*TreeNode)
                    // 注意取出栈顶之后，要将栈顶元素给清理
                    stack1 = stack1[:len(stack1)-1]
					res = append(res, cur.Val)
					cur = cur.Right
				} else {
					return res
				}
			}
        }
	}
    return res
}