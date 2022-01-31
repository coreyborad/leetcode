package no0206

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	// Make last node.Next = first.Next address
	var prev *ListNode
	cur := head
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	return prev
}
