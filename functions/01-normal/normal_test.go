package functions

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add Positive numbers", 2, 3, 5},
		{"Add negative numbers", -2, -3, -5},
		{"Add Positive & Negative Numbers", 10, -12, -2},
		{"Add Zeros", 0, 0, 0},
		{"Add Large Numbers", 1000000, 2000000, 3000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normal_add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("normal_add(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}

}
