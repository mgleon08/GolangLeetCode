# 387. First Unique Character in a String

[387. First Unique Character in a String
](https://leetcode.com/problems/first-unique-character-in-a-string/submissions/)

```go
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
```

# Version 1

```go
func firstUniqChar(s string) int {
	// 建立一個長度, 容量為 26 的 slice, counts[0 ~ 25]
	alphabet := make([]int, 26)

	for _, c := range s {
		// 單引號 '' 在 string 會轉化成 rune(int32) a=97, b=98..
		// 'a' - 'a' = 0, 'a' - 'b' = -1..
		alphabet[c-'a']++
	}

	for i, c := range s {
		if alphabet[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
```

Runtime: 8 ms
