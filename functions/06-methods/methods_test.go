package functions

import "testing"

func TestPerson_greet(t *testing.T) {
	tests := []struct {
		name   string
		person Person
		want   string
	}{
		{"Greet with John", Person{Name: "John"}, "Hello John"},
		{"Greet with Jane", Person{Name: "Jane"}, "Hello Jane"},
		{"Greet with empty name", Person{Name: ""}, "Hello "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.greet()
			if got != tt.want {
				t.Errorf("greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
