package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3}
	fmt.Println(m)

	d, ok := m["d"]
	fmt.Println(d, ok)

	m2 := map[string]string{
		"name": "afei",
		"age":  "18",
	}
	fmt.Println(m2)

	id, ok := m["id"]
	fmt.Println(id, ok)

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	m["id"] = 10
	fmt.Println(m)
}
