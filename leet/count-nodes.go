package leet

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	count := 0
	depth := calculateDepth(root)
	for root != nil {
		if calculateDepth(root.Right) == depth-1 {
			count += int(math.Pow(2, float64(depth)))
			root = root.Right
		} else {
			count += int(math.Pow(2, float64(depth)-1))
			root = root.Left
		}
		depth -= 1

	}
	return count
}

func calculateDepth(node *TreeNode) int {
	if node == nil {
		return -1
	}
	depth := 0
	for node.Left != nil {
		node = node.Left
		depth++
	}
	return depth
}
