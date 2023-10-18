package main

//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	if(p == nil && q == nil){
//		return true
//	}
//	if(p==nil || q==nil){
//		return false
//	}
//
//	if (p.Val != q.Val){
//		return false
//	}
//
//	lSame := isSameTree(p.Left, q.Left)
//	rSame := isSameTree(p.Right, q.Right)
//
//	return lSame && rSame
//
//}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	sameTree := false

	if root == nil {
		return false
	}

	if root.Val == subRoot.Val {
		sameTree = isSameTree(root, subRoot)
		if sameTree {
			return true
		}
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func main() {
	isSubtree(&TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	})
}
