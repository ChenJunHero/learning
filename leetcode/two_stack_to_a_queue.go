package main
/* 
栈是先进后出
队列先入先出
使用两个栈一个只负责插入，一个只负责删除
插入时：
直接insert到栈a的尾部
删除时：
b没有元素时，把a栈的所有元素出栈到b中
用于只删除b的栈顶元素
队列的所有元素为a+b
*/
type CQueue struct {
	stackA []int // a栈用户添加元素
	stackB []int // b栈用于删除元素
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	// 向a里面添加元素
	this.stackA = append(this.stackA, value)

}

func (this *CQueue) DeleteHead() int {
	if len(this.stackB) == 0 {
		// a,b 均为0， 则return -1
		if len(this.stackA) == 0 {
			return -1
		}
		// 将栈a所有元素出到b中
		for len(this.stackA) > 0 {
			index := len(this.stackA) - 1
			// 出栈a的栈顶元素到b中
			this.stackB = append(this.stackB, this.stackA[index])
			this.stackA = this.stackA[:index]
		}
	}

	// 删除栈b中最末尾的元素
	delIndex := len(this.stackB) - 1
	delValue := this.stackB[delIndex]
	this.stackB = this.stackB[:delIndex]
	return delValue
}

func main() {

}
