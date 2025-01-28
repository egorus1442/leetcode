package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	answer := head
	if head == nil {
		return nil
	}
	for answer.Next != nil {
		if answer.Val == answer.Next.Val {
			answer.Next = answer.Next.Next
		} else {
			answer = answer.Next
		}
	}
	return head
}
