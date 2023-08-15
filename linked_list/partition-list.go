package main

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	big := &ListNode{}
	_s, _b := small, big
	for head != nil {
		if head.Val < x {
			_s.Next = head
			_s = _s.Next
		} else {
			_b.Next = head
			_b = _b.Next
		}
		head = head.Next
	}
	_s.Next = big.Next
	_b.Next = nil
	return small.Next
}
