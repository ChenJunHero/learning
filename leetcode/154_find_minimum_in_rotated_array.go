package main

import "fmt"

/*
旋转数组中的最小值

思路：
二分查找，比较中间值和最右值
1.左右指针，分别指向首尾元素
2.循环求出中间指针
3.中间元素比右边元素小，指针右移到中间位置
4.中间元素比右边元素大，指针左移到中间位置+1
5.中间元素等于右边元素，右边指针减1
6.左边元素不小于右边元素退出即可
*/

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		mid := (left+right)/2
		switch {
		case numbers[mid] > numbers[right]:
			left = mid + 1
		case numbers[mid] < numbers[right]:
			right = mid
		case numbers[mid] == numbers[right]:
			right --
		}
	}
	return numbers[left]
}

func main() {
	test := []int{1,3,3}
	fmt.Println(minArray(test))

	a := "c"
	fmt.Println(byte(a), []byte(a)[0])
}
