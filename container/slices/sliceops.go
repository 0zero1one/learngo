package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr1)

	// 去除某个元素
	arr1 = append(arr1[:3], arr1[4:]...)
	fmt.Println(arr1)

	arr2 := make([]int, 10, 15)
	fmt.Println(arr2)
	fmt.Println(len(arr2), cap(arr2))
}
