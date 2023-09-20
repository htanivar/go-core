package sample

import (
	"fmt"
	"testing"
)

func ProcessData(data string, f func(string2 string) string) string {
	result := f(data)
	return result
}

func TestProcessData(t *testing.T) {

	mockFunc := func(data string) string {
		return "mocked result"
	}

	result := ProcessData("input data", mockFunc)
	fmt.Println(result)
}
