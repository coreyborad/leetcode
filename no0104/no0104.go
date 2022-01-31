package no0104

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// go until nil, only get deepest side tree
	t := max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	return t
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
