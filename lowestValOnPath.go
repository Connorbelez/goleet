package main

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestValOnPath(root *TreeNode, target *TreeNode, path map[int]*TreeNode) map[int]*TreeNode {
	if root == nil {
		return path
	}
	path[root.Val] = root
	if root.Val == target.Val {
		return path
	}
	if root.Val < target.Val {
		lowestValOnPath(root.Left, target, path)
	} else {
		lowestValOnPath(root.Right, target, path)
	}
	return path
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var pPath = make(map[int]*TreeNode)
	var qPath = make(map[int]*TreeNode)
	pPath = lowestValOnPath(root, p, pPath)
	qPath = lowestValOnPath(root, q, qPath)

	if len(qPath) < len(pPath) {
		pPath, qPath = qPath, pPath
	}
	lowestVal := math.MaxInt
	for key, _ := range pPath {
		_, ok := qPath[key]
		if ok {
			if lowestVal > key {
				lowestVal = key
			}
			continue
		}
	}

	//Find lowest val
	return pPath[lowestVal]

}
func main() {

}
