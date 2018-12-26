package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	// 透過 mark 去註記每次記錄到哪個 index
	mark := 0
	counts := 1
	for i := 0; i < len(chars); i++ {
		if i+1 != len(chars) && chars[i] == chars[i+1] {
			counts++
		} else {
			chars[mark] = chars[i]
			mark++
			if counts > 1 {
				// 將 counts 轉成 string 在個別塞到原本的 array 裡 e.g. 12 -> 1, 2
				for _, value := range []byte(strconv.Itoa(counts)) {
					chars[mark] = value
					mark++
				}
				counts = 1
			}
		}
	}
	// 最後只截取到 mark 的長度
	return len(chars[:mark])
}

func main() {
	testCase := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(testCase))
	testCase = []byte{'a'}
	fmt.Println(compress(testCase))
	testCase = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(compress(testCase))
}
