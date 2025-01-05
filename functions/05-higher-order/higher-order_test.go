package functions

import "testing"

func Test_applyFunction(t *testing.T) {
	tests := []struct {
		name  string
		a, b  int
		fname func(int, int) int
		want  int
	}{
		{"Add 2 and 3", 2, 3, add, 5},
		{"Minus 2 and 3", 2, 3, minus, -1},
		{"Times 2 and 3", 2, 3, times, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyFunction(tt.a, tt.b, tt.fname); got != tt.want {
				t.Errorf("applyFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
