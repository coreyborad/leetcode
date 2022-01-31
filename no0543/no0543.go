package no0543

import "github.com/halfrost/LeetCode-Go/structures"

// TreeNode define
type TreeNode = structures.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	L := dfs(root.Left, ans)
	R := dfs(root.Right, ans)

	*ans = max(*ans, L+R)
	return max(L, R) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
