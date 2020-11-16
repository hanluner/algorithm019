/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func preorder(root *Node) []int {
    if root == nil {
        return []int{}
    }
    var result []int
    
    result = append(result, root.Val)
    for _, r := range root.Children {
        result = append(result, preorder(r)...)
    }

    return result
}