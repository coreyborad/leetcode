package no0019

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	for ; n > 0; n-- {
		if head != nil {
			head = head.Next
		}
	}

	prev := dummy
	cur := dummy.Next
	for ; head != nil; head = head.Next {
		prev = prev.Next
		cur = cur.Next
	}
	prev.Next = cur.Next
	return dummy.Next
}
