# 73. Set Matrix Zeroes

[73. Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/)

```ruby
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

Example 1:

Input: 
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
Output: 
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
Example 2:

Input: 
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
Output: 
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
Follow up:

A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
```

這題主要是在考空間複雜度的部分，因此原本做法和參考做法時間複雜度都是 `O(MxN)` 

但在空間複雜度

* 原本做法: O(M+N)
    * 簡單地用兩個 array 紀錄，哪行 row 或是 col 要改成 0
* 參考做法: O(1)
    * 只要遇到 0 就將 row 和 col 的第一個值，改成 0，最後再判斷只要第一個是 0 就將整行改成 0

# Version 1

原本做法

```go
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
```

Runtime: 28 ms

# Version 2

參考做法

```go
package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {

	isCol := false
	rowLen := len(matrix)
	colLen := len(matrix[0])

	for i := 0; i < rowLen; i++ {
		//  用來判斷第一行的 col 有沒有 0，因為過程中會將 row 和 col 的第一個值更改為 0，之後就無法判斷
		if matrix[i][0] == 0 {
			isCol = true
		}

		// 接著都從 index 1 開始跑，只要有 0 就將 row 和 col 的第一個值更改為 0
		for j := 1; j < colLen; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 接著透過 row 和 col 的第一個值，來判斷是否將整行改為 0
	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 判斷第一個是不是 0，將第一行 row 都改成 0
	if matrix[0][0] == 0 {
		for j := 0; j < colLen; j++ {
			matrix[0][j] = 0
		}
	}
	// 判斷一開始設定的 isCol 來判斷是否第一行 col 都改成 0
	if isCol {
		for i := 0; i < rowLen; i++ {
			matrix[i][0] = 0
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
```

Runtime: 28 ms
