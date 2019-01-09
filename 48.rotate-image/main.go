package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	// 不需包含最後一位，所以減 1
	length := len(matrix) - 1
	// 只需要跑一邊，跑四個邊，就可以對調，因此先除 2
	for i := 0; i < len(matrix)/2; i++ {
		// 跑內圈
		for j := i; j < length-i; j++ {
			// 取得一開始的值
			tmp := matrix[i][j]
			// 定義一開始的位置
			nextX := j
			nextY := length - i
			// 跑四個邊
			for loop := 0; loop < 4; loop++ {
				// 和 tmp 值交換
				tmp, matrix[nextX][nextY] = matrix[nextX][nextY], tmp
				// 定義下一個位置
				tmpX := nextY
				nextY = length - nextX
				nextX = tmpX
			}

		}
	}

}

func main() {

	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}

	rotate(matrix)
	fmt.Println(matrix)
}
