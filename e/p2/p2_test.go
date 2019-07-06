package main

import (
	"testing"
)

func Test_fibonacci_evens_sum (t *testing.T) {
	tests := []struct {
		input, output int
	}{
		{ 5, 2 },
		{ 10, 10 },
		{ 40, 44 },
	}

	for _, test := range tests {
		got:=fibonacci_evens_sum(test.input)
		if got!=test.output {
			t.Errorf("Limit %d, got %d but wanted %d",test.input,got,test.output)
		}
	}

}