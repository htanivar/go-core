package functions

import "testing"

func Test_variadic_add(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Add Positive numbers", []int{2, 3}, 5},
		{"Add negative numbers", []int{-2, -3}, -5},
		{"Add Positive & Negative Numbers", []int{10, -12}, -2},
		{"Add Zeros", []int{0, 0}, 0},
		{"Add Large Numbers", []int{1000000, 2000000}, 3000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := variadic_add(tt.nums...); got != tt.want {
				t.Errorf("variadic_add() = %v, want %v", got, tt.want)
			}
		})
	}
}
