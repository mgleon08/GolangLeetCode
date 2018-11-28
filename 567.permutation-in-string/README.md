# 567. Permutation in String

* [567. Permutation in String
](https://leetcode.com/problems/permutation-in-string/)
* [Sliding Window Protocol](http://ecomputernotes.com/computernetworkingnotes/communication-networks/sliding-window-protocol)

```golang
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.

Example 1:
Input:s1 = "ab" s2 = "eidbaooo"
Output:True
Explanation: s2 contains one permutation of s1 ("ba").

Example 2:
Input:s1= "ab" s2 = "eidboaoo"
Output: False

Note:
The input strings only contain lower case letters.
The length of both given strings is in range [1, 10,000].
```

# Version 1

```golang
func checkInclusion(s1, s2 string) bool {
	mapS1 := slidingWindow(s1)

	for i := 0; i <= (len(s2) - len(s1)); i++ {
		// move s2 slidingWindow
		end := len(s1) + i
		mapS2 := slidingWindow(s2[i:end])
		// or use reflect.DeepEqual
		if intArrayEquals(mapS1, mapS2) {
			return true
		}
	}
	return false
}

func slidingWindow(chars string) []int {
	mapChar := make([]int, 26)
	for _, char := range chars {
		mapChar[char-'a']++
	}
	return mapChar
}

func intArrayEquals(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}
```

Runtime: 44 ms
