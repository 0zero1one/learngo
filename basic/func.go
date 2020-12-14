package main

import "fmt"

func div(a, b int) (int, int) {
	return a / b, a % b
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(div(5, 2))
	fmt.Println(swap(3, 4))
}
