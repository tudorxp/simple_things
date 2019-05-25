package main

import (
	"testing"
	// "fmt"
)

func TestFizzBuzz (t *testing.T) {
	tests := []struct{ 
		input int
		want string	
	}{
		{1, "1 "},
		{15,"1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz "},
	}

	for _, test := range tests {
		got := FizzBuzz(test.input)
		if got != test.want {
			t.Errorf("failed, want: %v, got: %v",test.want,got)
		}
	}

}