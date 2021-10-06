package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// root为空，公共节点为空
	if root == nil {
		return nil
	}
	// p,q有任何一个等于root，则直接返回root
	if root == p || root == q {
		return root
	}

	// 二叉树的后序遍历
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 情况1， 如果p和q都在以root为根的树中， left和right一定分别是p和q
	if left != nil && right != nil {
		return root
	}

	// 情况2， 如果p和q都不在以root为跟的节点中，返回nil
	if left == nil && right == nil {
		return nil
	}

	//  情况3， 如果p和q只有一个在root为跟的树中，函数返回该节点
	if left == nil {
		return right
	} else {
		return left
	}

}
func main() {
	fmt.Println("lca...")
}
