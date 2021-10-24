package main

import "fmt"
/* 

思路：
三个队列
1.循环终止条件
队列为空的时候

2.取数的逻辑：
每次取当前列表的前两个元素
如果当前元素不够取，则从下个列表的第一个元素取，"之"字形，直到满足m个队列
取完后，当前列表放到尾部


*/

func groupNlists(N [][]int, M int) [][]int {
	var maxLen int
	for _, n := range N {
		if len(n) >= maxLen {
			maxLen = len(n)
		}
	}
	var round int
	if maxLen%M == 0 {
		round = maxLen / M
	} else {
		round = maxLen/M + 1
	}
	fmt.Println("round", round, maxLen)
	var res [][]int
	for r := 0; r < round; r++ {
		e := getPopElems(r, N, M)
		res = append(res, e...)
	}
	return res
}

func getPopElems(round int, N [][]int, M int) [][]int {
	var popelems []int
	for i := range N {
		if len(N[i]) == 0 {
			continue
		}
		if len(N[i]) <= M {
			popelems = append(popelems, N[i]...)
			N[i] = []int{}
		} else {
			popelems = append(popelems, N[i][:M]...)
			N[i] = N[i][M:]
		}
	}
	var res [][]int
	for _ = range popelems {
		if len(popelems) <= M {
			res = append(res, popelems)
			break
		}
		res = append(res, popelems[:M])
		popelems = popelems[M:]
	}
	fmt.Println(round, res)
	return res
}

func dividNums(nums [][]int, m int) [][]int {
	if len(nums) <= 0 || m <= 0 {
		return [][]int{}
	}

	res := [][]int{}
	queue := nums
	for len(queue) > 0 {
		// curr := []int{}
		curr := make([]int, 0, m)
		for len(queue) > 0 {
			//
			front := queue[0]
			for len(curr) < m && len(front) > 0 {
				curr = append(curr, front[0])
				front = front[1:]
			}
			// pop up
			queue = queue[1:]
			if len(front) > 0 {
				queue = append(queue, front)
			}
			if len(curr) == m {
				break
			}
		}
		res = append(res, curr)
	}
	return res
}

func main() {
	N := [][]int{
		{1, 2, 3, 4, 5, 6},
		{11, 12, 13},
		{21, 22, 23, 24},
	}
	M := 2
	// res := groupNlists(N, M)
	// fmt.Println(res)
	fmt.Println(dividNums(N, M))
}
