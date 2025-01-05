package functions

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func applyFunction(a, b int, fname func(int, int) int) int {
	return fname(a, b)
}
