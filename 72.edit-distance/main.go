package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	// 如果其中一個長度是0，最小距離就會是兩個加起來
	if len1*len2 == 0 {
		return len1 + len2
	}
	// 組一個二維陣列，用於記錄每次最小的步驟
	// +1 是考慮到DP中，一個串為空的情況
	arr2d := make([][]int, len1+1)
	for i := range arr2d {
		arr2d[i] = make([]int, len2+1)
	}
	// 初始化 Array
	for i := 1; i < len1+1; i++ {
		arr2d[i][0] = i
	}
	for j := 1; j < len2+1; j++ {
		arr2d[0][j] = j
	}
	// Dynamic Programming 計算
	numbers := make([]int, 3)
	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			// 左邊和下面因為少一個字，所以一定會多一個步驟
			numbers[0] = arr2d[i-1][j] + 1
			numbers[1] = arr2d[i][j-1] + 1
			// 中間的字，如果不相等，那就要多一個步驟才能變成一樣
			numbers[2] = arr2d[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				numbers[2]++
			}
			//  將裡面最小值，填入新的欄位
			arr2d[i][j] = minNumber(numbers)
		}
	}
	// 最後面一個就是可以組成兩個字相等時所需的最小步驟
	return arr2d[len1][len2]
}

func minNumber(numbers []int) int {
	min := numbers[0]
	for _, value := range numbers {
		if value < min {
			min = value
		}
	}
	return min
}

func main() {
	word1, word2 := "horse", "ros"
	fmt.Println(minDistance(word1, word2))
	word1, word2 = "intention", "execution"
	fmt.Println(minDistance(word1, word2))
}
