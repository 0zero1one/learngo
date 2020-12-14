package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	a := [10]int{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i] ++
				//fmt.Printf("hello from goroutine %d\n", i)
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
