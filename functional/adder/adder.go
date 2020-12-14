package adder

func Adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type IAdder func(int) (int, IAdder)

func StandardAdder(base int) IAdder {
	return func(v int) (int, IAdder) {
		return base + v, StandardAdder(base + v)
	}
}
