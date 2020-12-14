package main

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func RecFib(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	return RecFib(number - 1) + RecFib(number - 2)
}

func CountFib(a, b, count int) int {
	if count == 0 {
		return b
	}

	return CountFib(b, a + b, count - 1)
}