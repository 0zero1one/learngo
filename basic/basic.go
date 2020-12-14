package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 2, 4
	var s string = "123"
	fmt.Println(a , b, s)
	fmt.Println(i)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func triangle() {
	a, b := 3, 4
	fmt.Println(math.Sqrt(float64(a * a + b * b)))
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	c := math.Sqrt(a * a + b * b)
	fmt.Println(filename, c)
}

func enmus() {
	const(
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(a, b, c)
	const(
		d = iota
		e
		f
	)
	fmt.Println(d, e, f)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	fmt.Println(i)
	euler()
	triangle()
	consts()
	enmus()
}

var i = 1
