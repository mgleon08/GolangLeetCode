# 72. Edit Distance

[72. Edit Distance
](https://leetcode.com/problems/edit-distance/)

```go
Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
```

# Version 1

這是一個經典的演算法，由俄羅斯科學家Vladimir Levenshtein在1965年提出這個概念。

這題主要是用到了，DP(Dynamic Programming) 這個概念來解

* [72. Edit Distance](https://leetcode.com/articles/edit-distance/)
* [編輯距離演算法（Levenshtein）](https://tw.saowen.com/a/633d345525ad4da6f65d03e172777ee1749e8cbce92c3b225ae63c123ea61ebb)
* [[理工] 演算法 min edit請教](https://www.ptt.cc/bbs/Grad-ProbAsk/M.1470239320.A.268.html)
* [動態規劃(DP)的整理-Python描述](https://blog.csdn.net/MrLevo520/article/details/75676160)
* [Dynamic Programming](http://www.csie.ntnu.edu.tw/~u91029/DynamicProgramming.html)

```go
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
```

Runtime: 12 ms
