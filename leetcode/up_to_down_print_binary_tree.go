package main

import "fmt"

/*
剑指 Offer 32 - I. 从上到下打印二叉树
BFS实现二叉树的广度优先搜索

算法流程：
1.root节点入队列
2.出対添加到打印结果
3.把左右节点加到队列中
4.队列为空中止循环
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		// 队首元素出对
		curNode := queue[0]
		queue = queue[1:]
		res = append(res, curNode.Val)
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}
	return res
}

func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n1.Left = n2
	fmt.Println(levelOrder(n1))
}
