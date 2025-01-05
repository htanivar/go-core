package functions

import "math"

func normal_add(a, b int) int {
	return a + b
}

func normal_minus(a, b int) int {
	return a - b
}

func normal_times(a, b int) int {
	return a * b
}

func normal_division(a, b float64) float64 {
	if a == 0 && b == 0 {
		return 0
	}
	result := a / b
	roundedResult := math.Round(result*100) / 100
	return roundedResult
}
