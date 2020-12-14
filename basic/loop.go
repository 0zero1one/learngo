package main

import (
	"bufio"
	"fmt"
	"os"
)

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println("for")
	printFile("abc.txt")
}
