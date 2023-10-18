package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	h := 1
	left, right := 0, 0
	lb, rb := true, true

	left, lb = balanced(root.Left)
	right, rb = balanced(root.Right)

	h = h + int(math.Max(float64(left), float64(right)))
	if !(rb && lb) {
		return -1, false
	}

	return h, math.Abs(float64(left-right)) <= 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, b := balanced(root)

	return b
}

func main() {
	isBalanced(&TreeNode{Val: 10})
}
