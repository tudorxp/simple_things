package main

import (
	"testing"
)

func Test_is_palindrome (t *testing.T) {
	tests := []struct{
		input int
		want bool
	}{
		{ 12315, false },
		{ 12317823, false },
		{ 128313821, true },
		{ 4, true },
		{ 44, true },
		{ 8671, false },
	}

	for _, test := range tests {
		got := is_palindrome(test.input)
		if got != test.want {
			t.Errorf("Error for test: %d, want: %v but got %v",test.input,test.want,got)
		}
	}
}