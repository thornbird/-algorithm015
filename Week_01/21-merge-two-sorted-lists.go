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

    var mergeList *ListNode

    if l1.Val <= l2.Val {
        mergeList = l1
        mergeList.Next = mergeTwoLists(l1.Next, l2)
        l1 = l1.Next
    } else {
        mergeList = l2
        mergeList.Next = mergeTwoLists(l2.Next, l1)
        l2 = l2.Next
    }

    return mergeList

}