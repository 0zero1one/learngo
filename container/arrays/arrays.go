package main

import "fmt"

func main() {
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr1)

	s1 := arr1[1:4]
	fmt.Println(s1)

	s1[0] = 10
	fmt.Println(arr1)
	fmt.Println(s1)

	s2 := s1[1:]
	fmt.Println(s2)

	s2[0] = 10
	fmt.Println(s2)
	fmt.Println(s1)
	fmt.Println(arr1)

	s3 := s2[1:4]
	fmt.Println(s3)

	fmt.Println(cap(s3))
	fmt.Println(len(s3))

	s4 := append(s3, 9)
	s5 := append(s4, 10)
	s6 := append(s5, 11)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(arr1)
	fmt.Println(len(arr1), cap(arr1))
}
