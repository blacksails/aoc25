package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet0Visits(t *testing.T) {
	tests := []struct {
		pos      int
		steps    int
		expected int
	}{
		{pos: 50, steps: 60, expected: 1},
		{pos: 90, steps: 20, expected: 1},
		{pos: 0, steps: 100, expected: 1},
		{pos: 99, steps: 1, expected: 1},
		{pos: 10, steps: 5, expected: 0},
		{pos: 95, steps: 10, expected: 1},
		{pos: 50, steps: 150, expected: 2},
		{pos: 50, steps: -150, expected: 2},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := get0Visits(test.pos, test.steps)
			assert.Equal(t, test.expected, actual)
		})
	}
}
