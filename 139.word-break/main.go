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
