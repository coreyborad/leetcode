package no0141

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// slow is 1 step, fast is 2 step
	slow := head
	fast := head.Next

	for {
		if fast == nil || fast.Next == nil {
			break
		}
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
