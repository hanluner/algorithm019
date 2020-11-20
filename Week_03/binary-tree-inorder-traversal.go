/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int

	if root.Left != nil {
		leftSlice := inorderTraversal(root.Left)
		result = append(result, leftSlice...)
	}

	result = append(result, root.Val)

	if root.Right != nil {
		rightSlice := inorderTraversal(root.Right)
		result = append(result, rightSlice...)
	}

	return result
}