package functions

func anonymous_add(a, b int) int {
	add := func(a, b int) int {
		return a + b
	}

	return add(a, b)
}
