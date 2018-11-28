package main

import "fmt"

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

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
	s1 = "ab"
	s2 = "eidboaoo"
	fmt.Println(checkInclusion(s1, s2))

	s1 = "prosperity"
	s2 = "properties"
	fmt.Println(checkInclusion(s1, s2))

	s1 = "a"
	s2 = "ab"
	fmt.Println(checkInclusion(s1, s2))

	s1 = "adc"
	s2 = "dcda"
	fmt.Println(checkInclusion(s1, s2))

	s1 = "abcdxabcde"
	s2 = "abcdeabcdx"
	fmt.Println(checkInclusion(s1, s2))
}
