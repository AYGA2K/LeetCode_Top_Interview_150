package leet

import (
	"math"
)

/*
*
  - Definition for a binary tree node.
    type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
    }
*/

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDiff := math.MaxInt64
	var prev *TreeNode
	dfs(root, &prev, &minDiff)
	return minDiff
}

func dfs(node *TreeNode, prev **TreeNode, minDiff *int) {
	if node == nil {
		return
	}
	dfs(node.Left, prev, minDiff)
	if *prev != nil {
		diff := int(math.Abs(float64(node.Val - (*prev).Val)))
		if diff < *minDiff {
			*minDiff = diff
		}
	}
	*prev = node
	dfs(node.Right, prev, minDiff)
}
