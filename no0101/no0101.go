package no0101

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	reverseNode := reverseTree(root.Right)
	return sameTree(root.Left, reverseNode)
}

func reverseTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left, node.Right = node.Right, node.Left
	reverseTree(node.Left)
	reverseTree(node.Right)
	return node
}

func sameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q == nil {
		return false
	}

	if p == nil && q != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !sameTree(p.Left, q.Left) {
		return false
	}

	if !sameTree(p.Right, q.Right) {
		return false
	}

	return true
}
