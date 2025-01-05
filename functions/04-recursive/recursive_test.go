package functions

import "testing"

func Test_factorial(t *testing.T) {
	tests := []struct {
		name     string
		x        uint
		expected uint
	}{
		{"Factorial of 0", 0, 1},
		{"Factorial of 1", 1, 1},
		{"Factorial of 2", 2, 2},
		{"Factorial of 3", 3, 6},
		{"Factorial of 4", 4, 24},
		{"Factorial of 5", 5, 120},
		{"Factorial of 6", 6, 720},
		{"Factorial of 7", 7, 5040},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := factorial(tt.x)
			if got != tt.expected {
				t.Errorf("factorial(%d) = %d; want %d", tt.x, got, tt.expected)
			}
		})
	}
}
