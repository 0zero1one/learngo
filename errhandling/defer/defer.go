package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 06666)
	defer file.Close()

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	writeFile("fib.txt")
}
