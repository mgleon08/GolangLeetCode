package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	rowLen := len(matrix)
	colLen := len(matrix[0])
	// 建立 col 和 row 的 map 來記錄
	rowMap := make([]int, rowLen)
	colMap := make([]int, colLen)

	// 只要有 0 就記錄在 rowMap && colMap
	for i := 0; i < rowLen; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowMap[i] = 1
				colMap[j] = 1
			}

		}
	}

	// 接著只要 rowMap 或是 colMap 有 1 就改成 0
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if rowMap[i] == 1 || colMap[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}

}

func main() {
	matrix := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}

	setZeroes(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}

	setZeroes(matrix)
	fmt.Println(matrix)
}
