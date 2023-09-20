package testy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//
//func TestCalculate(t *testing.T) {
//	assert.Equal(t, Calculate(2), 4)
//}

func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{input: 2, expected: 4},
		{input: -1, expected: 1},
		{input: 0, expected: 2},
		{input: -5, expected: -3},
		{input: 99999, expected: 100001},
	}
	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}
