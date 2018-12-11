# 139. Word Break

[139. Word Break](https://leetcode.com/problems/word-break//)

```go
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
```

* [Backtracking](http://www.csie.ntnu.edu.tw/~u91029/Backtracking.html)
* [Graph: Depth-First Search(DFS，深度優先搜尋)](http://alrightchiu.github.io/SecondRound/graph-depth-first-searchdfsshen-du-you-xian-sou-xun.html)
* [深度優先搜尋法 (Depth-first Search)](http://simonsays-tw.com/web/DFS-BFS/DepthFirstSearch.html)
* [Dynamic Programming](http://www.csie.ntnu.edu.tw/~u91029/DynamicProgramming.html)
* [What's the difference between backtracking and depth first search?](https://stackoverflow.com/questions/1294720/whats-the-difference-between-backtracking-and-depth-first-search)

# Version 1

```go
package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	markArray := make([]bool, len(s))
	return isMatched(s, 0, wordDict, &markArray)
}

func isMatched(s string, start int, wordDict []string, markArray *[]bool) bool {
	// 如果最後一個已經 matched，就不用繼續下去，直接 return
	if (*markArray)[len(s)-1] {
		return true
	}
	for i := start; i < len(s); i++ {
		// 有 matchd 過就繼續下一個
		if (*markArray)[i] {
			continue
		}
		if checkDict(s[start:i+1], wordDict) {
			// 如果有 matched 就註記
			(*markArray)[i] = true
			if isMatched(s, i+1, wordDict, markArray) {
				return true
			}
		}
	}
	return false
}

func checkDict(s string, wordDict []string) bool {
	for _, v := range wordDict {
		if s == v {
			return true
		}
	}
	return false
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))

	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDict))

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))

	s = "aaaaaaa"
	wordDict = []string{"aaaa", "aaa"}
	fmt.Println(wordBreak(s, wordDict))
}
```

Runtime: 0 ms
