package no0617

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var newTree *TreeNode
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 != nil && root2 == nil {
		newTree = &TreeNode{
			Val:   root1.Val,
			Left:  mergeTrees(root1.Left, nil),
			Right: mergeTrees(root1.Right, nil),
		}
	} else if root1 == nil && root2 != nil {
		newTree = &TreeNode{
			Val:   root2.Val,
			Left:  mergeTrees(nil, root2.Left),
			Right: mergeTrees(nil, root2.Right),
		}
	} else {
		newTree = &TreeNode{
			Val:   root1.Val + root2.Val,
			Left:  mergeTrees(root1.Left, root2.Left),
			Right: mergeTrees(root1.Right, root2.Right),
		}
	}

	return newTree
}
