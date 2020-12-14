package main

import (
	"fmt"
	"io/ioutil"
)

func grage(score int) string {
	var g string
	switch {
	case score < 0 || score > 100:
		panic("score in [0, 100]")
	case score <= 60:
		g = "D"
	case score <= 70:
		g = "C"
	case score <= 80:
		g = "B"
	case score <= 90:
		g = "A"
	case score <= 100:
		g = "S"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(
		grage(60),
		grage(70),
		grage(80),
		grage(90),
		grage(100),
		grage(-1),
		)
}
