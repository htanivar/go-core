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

func TestMinus(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Minus negative numbers", -2, -3, 1},
		{"Minus Positive & Negative Numbers", 10, -12, 22},
		{"Minus Zeros", 0, 0, 0},
		{"Minus Large Numbers", 1000000, 2000000, -1000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normal_minus(tt.a, tt.b); got != tt.expected {
				t.Errorf("normal_minus() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTimes(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Times positive numbers", 2, 3, 6},
		{"Times negative numbers", -2, -3, 6},
		{"Times Positive & Negative Numbers", 10, -12, -120},
		{"Times Zeros", 0, 0, 0},
		{"Times Large Numbers", 1000000, 2000000, 2000000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normal_times(tt.a, tt.b); got != tt.want {
				t.Errorf("normal_times() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"Divide positive numbers", 2, 3, 0.67},
		{"Divide negative numbers", -2, -3, 0.67},
		{"Divide Positive & Negative Numbers", 10, -12, -0.83},
		{"Divide Zeros", 0, 0, 0},
		{"Divide Large Numbers", 1000000, 2000000, 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normal_division(tt.a, tt.b); got != tt.want {
				t.Errorf("normal_division() = %v, want %v", got, tt.want)
			}
		})
	}
}
