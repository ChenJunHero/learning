package main

import "fmt"

/*
在排序数组中查找数字 I
二分查找
*/

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		//  左边找
		if nums[mid] > target {
			right = mid - 1
			continue
		}

		//  右边找
		if nums[mid] < target {
			left = mid + 1
			continue
		}
	}
	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return 0
	}

	//  找到目标索引
	index := binarySearch(nums, target)
	if index == -1 {
		return 0
	}
	count := 0
	for i := index; i < len(nums); i++ {
		if nums[i] != target {
			break
		}
		count++
	}
	for i := index - 1; i >= 0; i-- {
		if nums[i] != target {
			break
		}
		count++
	}
	return count
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search(nums, 8))
}
