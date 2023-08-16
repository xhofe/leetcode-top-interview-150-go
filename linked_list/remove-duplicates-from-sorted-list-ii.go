package main

func deleteDuplicates(head *ListNode) *ListNode {
	val := -101
	var prev, newHead *ListNode
	for head != nil {
		if head.Val != val && (head.Next == nil || head.Next.Val != head.Val) {
			if prev == nil {
				newHead = head
			} else {
				prev.Next = head
			}
			prev = head
		}
		val = head.Val
		head = head.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return newHead
}
