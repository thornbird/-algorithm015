/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    fast, slow := head, head
    
    // 注意这里可以不用判断 slow 是否为 nil 因为fast一定会比slow先到达nil; 或者有环永远不会为nil
    // 这里也不用判断head 若head为nil, 则for循环不会进去，直接返回
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        
        if slow == fast {
            return true
        }
    }
    
    return false
}