package main

/*
方法1，暴力搜索
方法2: 线性搜索，减少搜索空间
从左上开始搜索，分为三种情况：
1. 等于当前值，返回true
2. 大于当前值，列下标-1
3. 小于当前值，行下标+1
时间复杂度O(n+m)
空间复杂度O(1)
*/

//  暴力搜索
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	// bound case
	if len(matrix) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == target {
				return true
			}
		}
	}
	return false
}

//  线性搜索
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	// bound case
	if len(matrix) == 0 {
		return false
	}
	var x, y int
	y = len(matrix[0]) - 1
	for {
		if y < 0 || x > len(matrix)-1 {
			return false
		}

		if matrix[x][y] == target {
			return true
		}
		// 当前值大于目标值，列数-1
		// 当前值小于目标值，行数+1
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
}
func main() {

}
