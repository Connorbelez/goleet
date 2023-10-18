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

func smallestValOnPath(root *TreeNode, target *TreeNode, path map[int]*TreeNode) map[int]*TreeNode {
	if root == nil {
		return path
	}
	path[root.Val] = root
	if root.Val == target.Val {
		return path
	}
	if root.Val < target.Val {
		smallestValOnPath(root.Left, target, path)
	} else {
		smallestValOnPath(root.Right, target, path)
	}
	return path
}

func smallestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var pPath = make(map[int]*TreeNode)
	var qPath = make(map[int]*TreeNode)
	pPath = smallestValOnPath(root, p, pPath)
	qPath = smallestValOnPath(root, q, qPath)

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
