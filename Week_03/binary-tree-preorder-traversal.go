/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	result = append(result, root.Val)

	if root.Left != nil {
		leftSlice := preorderTraversal(root.Left)
		result = append(result, leftSlice...)
	}

	if root.Right != nil {
		rightSlice := preorderTraversal(root.Right)
		result = append(result, rightSlice...)
	}

	return result

}