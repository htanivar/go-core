package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(5, 5)

	if result != 10 {
		t.Errorf("Add(5,5) FAILED. Expected %d, got %d \n", 10, result)
	} else {
		t.Logf("Add(5,5) PASSED. Expected %d, got %d \n", 10, result)
	}
}
