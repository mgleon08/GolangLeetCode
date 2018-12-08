package main

import "fmt"

func partition(s string) [][]string {
	result := [][]string{}
	list := make([]string, 0, len(s))
	// result 必須用 address，因為最後會將所有對應到的組合存取在同一個 result 當中
	dfs(s, list, &result)
	return result
}

func dfs(s string, list []string, result *[][]string) {
	if len(s) == 0 {
		tmp := make([]string, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
		return
	}
	for i := 1; i <= len(s); i++ {
		// 如果是 Palindrome 就將它加入 list, 並將後面的字段，繼續做判斷
		if isPalindrome(s[0:i]) {
			dfs(s[i:], append(list, s[0:i]), result)
		}
	}
}

func isPalindrome(s string) bool {
	// 由前後兩兩去對應，奇數因為中間的一個字一定是 Palindrome，因此不影響
	for i := 0; i < (len(s) / 2); i++ {
		if s[i] != s[len(s)-(i+1)] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("cdd"))
	fmt.Println(partition("eee"))
}
