package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestValOnPath(root *TreeNode, target *TreeNode, path []*TreeNode) []*TreeNode {
	if root == nil {
		return path
	}
	path = append(path, root)

	if root.Val == target.Val {
		return path
	}
	if root.Val > target.Val {
		path = lowestValOnPath(root.Left, target, path)
	} else {
		path = lowestValOnPath(root.Right, target, path)
	}
	return path
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := lowestValOnPath(root, p, []*TreeNode{})
	qPath := lowestValOnPath(root, q, []*TreeNode{})

	if len(qPath) < len(pPath) {
		pPath, qPath = qPath, pPath
	}

	for i := len(pPath) - 1; i >= 0; i-- {

		if pPath[i] == qPath[i] {
			return pPath[i]
		}
	}
	return nil

}
func main() {

}
