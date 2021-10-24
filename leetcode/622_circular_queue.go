package main
/* 
单链表实现
*/

type Node struct {
	Val int
	Next *Node
}

type MyCircularQueue struct {
	cap int 
	head *Node
	tail *Node
	count int
}


func Constructor(k int) MyCircularQueue {
	cq := MyCircularQueue{}
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Next = head
	return cq
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	// 判断队列已满
	if this.IsFull() {
		return false
	}
	

}


func (this *MyCircularQueue) DeQueue() bool {

}


func (this *MyCircularQueue) Front() int {

}


func (this *MyCircularQueue) Rear() int {

}


func (this *MyCircularQueue) IsEmpty() bool {

}


func (this *MyCircularQueue) IsFull() bool {
	if this.count >= this.cap {
		return true
	}
	return false
}