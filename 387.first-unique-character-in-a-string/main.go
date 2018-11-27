package main

import "fmt"

func firstUniqChar(s string) int {
	alphabet := make([]int, 26)
	for _, c := range s {
		alphabet[c-'a']++
	}
	for i, c := range s {
		if alphabet[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar(""))
	fmt.Println(firstUniqChar("cc"))
}
