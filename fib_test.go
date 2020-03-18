package fib

import (
	"testing"
	"fmt"
)

func TestFib(t *testing.T) {
	cases := []struct {index, value int}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{20, 6765},
		{30, 832040},
		{40, 102334155},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Fibonnaci number %d", c.index), func (t *testing.T) {
			calculated := Fib(c.index)
			if calculated != c.value {
				t.Fatalf("Failed calculating Fibonnaci number %d: expected %d, got %d", c.index, c.value, calculated)
			}
		})
	}
}
