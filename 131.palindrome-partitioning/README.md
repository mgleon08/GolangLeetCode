# 131. Palindrome Partitioning

[131. Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/)

```go
Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:

Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
```

* [Palindrome](http://www.csie.ntnu.edu.tw/~u91029/Palindrome.html)
* [Backtracking](http://www.csie.ntnu.edu.tw/~u91029/Backtracking.html)
* [Graph: Depth-First Search(DFS，深度優先搜尋)](http://alrightchiu.github.io/SecondRound/graph-depth-first-searchdfsshen-du-you-xian-sou-xun.html)
* [深度優先搜尋法 (Depth-first Search)](http://simonsays-tw.com/web/DFS-BFS/DepthFirstSearch.html)
* [Dynamic Programming](http://www.csie.ntnu.edu.tw/~u91029/DynamicProgramming.html)
* [What's the difference between backtracking and depth first search?](https://stackoverflow.com/questions/1294720/whats-the-difference-between-backtracking-and-depth-first-search)

# Version 1

```go
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
	// 只有前面都有對應到的，才會到最後面，否則再判斷 Palindrome 就會跳掉
	if len(s) == 0 {
		tmp := make([]string, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
		return
	}
	for i := 1; i <= len(s); i++ {
		// 如果是 Palindrome 就將它加入 list, 並將後面的字段，繼續做判斷
		if isPalindrome(s[0:i]) {
			// 每次都將對應之後的字組拿掉在傳入
			dfs(s[i:], append(list, s[0:i]), result)
		}
	}
}

func isPalindrome(s string) bool {
	// 前後兩兩去對應，奇數因為中間的一個字一定是 Palindrome，因此不影響
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
```

Runtime: 24 ms

# Version 2

```go
package main

import "fmt"

func partition(s string) [][]string {
	result := [][]string{}
	list := make([]string, 0, len(s))
	// result 必須用 address，因為最後會將所有對應到的組合存取在同一個 result 當中
	dfs(s, 0, list, &result)
	return result
}

func dfs(s string, i int, list []string, result *[][]string) {
	// 只有前面都有對應到的，才會到最後面，否則再判斷 Palindrome 就會跳掉
	if i == len(s) {
		tmp := make([]string, len(list))
		copy(tmp, list)
		*result = append(*result, tmp)
		return
	}
	for j := i; j < len(s); j++ {
		// 如果是 Palindrome 就將它加入 list, 並將後面的字段，繼續做判斷
		if isPalindrome(s[i:j+1]) {
			// 每次都將起始的 index + 1
			dfs(s, j+1, append(list, s[i:j+1]), result)
		}
	}
}

func isPalindrome(s string) bool {
	// 前後兩兩去對應，奇數因為中間的一個字一定是 Palindrome，因此不影響
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
```

Runtime: 24 ms

# Flow

### aab

```go
1.1 aab, a, list = [a], next = ab
  2.1 ab, a, list = [a a], next = b
    3.1 b, b, list = [a a b] return
  2.2 ab, ab, list = [a] false
1.2 aab, aa, list = [aa], next = b
  2.1 b, b, list = [aa b] return
1.3 aab, aab, list = [] false

// [[a a b] [aa b]]
```

### eee

```go
1.1 eee, e, list = [e], next = ee
  2.1 ee, e, list = [e e], next = e
    3.1 e, e, list = [e e e] return
  2.2 ee, ee, list = [e ee] return
1.2 eee, ee, list = [ee], next = e
  2.1 e, e, list = [ee e] return
1.3 eee, eee, list = [eee] return

// [[e e e] [e ee] [ee e] [eee]]
```
