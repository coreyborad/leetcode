package no0160

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB

	// run until nil by each with virtual pointer
	for curA != nil && curB != nil {
		curA = curA.Next
		curB = curB.Next
	}

	// run adjust point for real head
	for curA != nil {
		headA = headA.Next
		curA = curA.Next
	}

	for curB != nil {
		headB = headB.Next
		curB = curB.Next
	}

	// Becasue they are same position now, you can just compare A and B until they different
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	// Whatever you want to retrun A or B
	return headA
}
