package no0094

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) []int {
	// left=>root=>right
	if root == nil {
		return []int{}
	}
	left := []int{}
	right := []int{}

	if root.Left != nil {
		left = inorderTraversal(root.Left)
	}

	left = append(left, root.Val)

	if root.Right != nil {
		right = inorderTraversal(root.Right)
	}

	// left first and append right each
	return append(left, right...)
}
