package math

func Add(x, y int) (res int) {
	return x + y
}

func Multiply(x, y int) int {
	result := 0
	for i := y; i > 0; i-- {
		result = result + x
	}
	return result
}
