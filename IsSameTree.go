package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	lSame := isSameTree(p.Left, q.Left)
	rSame := isSameTree(p.Right, q.Right)

	return lSame && rSame

}
func main() {
	isSameTree(&TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	})

}
