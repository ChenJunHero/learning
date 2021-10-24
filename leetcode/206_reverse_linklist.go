package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代解法
func reverseList(head *ListNode) *ListNode {
	cur := head
	var res *ListNode
	for cur != nil {
		// 首先存储下一个节点的引用
		next := cur.Next
		// 当前节点指向反转后的链表，调整指针指向
		cur.Next = res
		// 反转后的链表头部等于当前节点
		res = cur
		// 遍历下一节点
		cur = next
	}
	return res
}

// 递归写法
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	a.Next = b
	b.Next = c
	
	reverse := reverseList(a)
	for reverse != nil {
		fmt.Println(reverse)
		reverse = reverse.Next
	}
}

