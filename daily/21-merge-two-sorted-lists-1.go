/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	l1Cur := l1
	l2Cur := l2

	var mergeList *ListNode

	// 初始化
	if l1Cur.Val <= l2Cur.Val {
		mergeList = l1Cur
		l1Cur = l1Cur.Next
	} else {
		mergeList = l2Cur
		l2Cur = l2Cur.Next
	}

	var tmp *ListNode = mergeList

	for true {
		if l1Cur == nil {
			tmp.Next = l2Cur
			break
		}
		if l2Cur == nil {
			tmp.Next = l1Cur
			break
		}

		if l1Cur.Val <= l2Cur.Val {
			tmp.Next = l1Cur
			l1Cur = l1Cur.Next
		} else {
			tmp.Next = l2Cur
			l2Cur = l2Cur.Next
		}
		tmp = tmp.Next
	}
	return mergeList
}