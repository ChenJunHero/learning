package main
/* 
剑指 Offer 30. 包含min函数的栈
最小栈问题：
可以方便的随时获取一个栈的最小元素

思路：
两个栈，一个用来实现常规的进栈，出栈，和top操作
一个辅助栈，把最小的元素放进去
*/
import "fmt"

type MinStack struct {
	stackA []int // 维护进栈，出栈
	stackB []int // 用于存储降序元素
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func top(stack []int) int {
	return stack[len(stack)-1]
}

func (this *MinStack) Push(x int) {
	this.stackA = append(this.stackA, x)
	// 入栈元素小于等于辅助栈顶元素, 也入栈
	if  len(this.stackB) == 0 || x <= top(this.stackB) {
		this.stackB = append(this.stackB, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stackA) == 0 {
		return
	}
	indexA := len(this.stackA) - 1
	popValue := this.stackA[indexA]
	this.stackA = this.stackA[:len(this.stackA)-1]
	if popValue == top(this.stackB) {
		this.stackB = this.stackB[:len(this.stackB)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stackA[len(this.stackA)-1]
}

func (this *MinStack) Min() int {
	if len(this.stackB) == 0 {
		return 0
	}
	return this.stackB[len(this.stackB)-1]
}

func main() {
	obj := Constructor()
	push := []int{2, 0, 3,0}
	for _, t := range push {
		obj.Push(t)
	}
	fmt.Println(obj.stackA, obj.stackB)
	for i:=0; i<3; i++ {
		obj.Pop()
	}
	fmt.Println(obj.Min())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
