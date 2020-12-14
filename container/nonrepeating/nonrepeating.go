package main

import (
	"fmt"
)

func norepeating(s string) int {
	lastO := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastO[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastO[ch] = i
	}
	return maxLength
}

func mbnorepeating(s string) int {
	lastO := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastO[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastO[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println("norepeating:")
	fmt.Println(norepeating("abcabab"))
	fmt.Println(norepeating("123456123456712"))

	fmt.Println("mbnorepeating:")
	fmt.Println(mbnorepeating("abcabab"))
	fmt.Println(mbnorepeating("123456123456712"))
	fmt.Println(mbnorepeating("我是阿飞,你是谁啊"))
	fmt.Println(mbnorepeating("我是afei,你是谁啊"))
}
