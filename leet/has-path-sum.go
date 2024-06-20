package leet

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return traverse(root, 0, targetSum)
}
func traverse(node *TreeNode, sum, targetSum int) bool {
	if node == nil {
		return false
	}
	sum = sum + node.Val
	if node.Left == nil && node.Right == nil {
		return sum == targetSum
	}
	return traverse(node.Left, sum, targetSum) || traverse(node.Right, sum, targetSum)
}
