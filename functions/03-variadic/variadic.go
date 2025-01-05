package functions

func variadic_add(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
