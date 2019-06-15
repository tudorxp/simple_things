package main

import "testing"

func TestPascal(t *testing.T) {
	type testcase struct {
		input int
		want  string
	}
	tests := []testcase{
		{4, "[1]\n[1 1]\n[1 2 1]\n[1 3 3 1]\n"},
		{8, "[1]\n[1 1]\n[1 2 1]\n[1 3 3 1]\n[1 4 6 4 1]\n[1 5 10 10 5 1]\n[1 6 15 20 15 6 1]\n[1 7 21 35 35 21 7 1]\n"},
		{10, `[1]
[1 1]
[1 2 1]
[1 3 3 1]
[1 4 6 4 1]
[1 5 10 10 5 1]
[1 6 15 20 15 6 1]
[1 7 21 35 35 21 7 1]
[1 8 28 56 70 56 28 8 1]
[1 9 36 84 126 126 84 36 9 1]
`},
	}

	for _, test := range tests {
		output:=pascal(test.input)
		if output != test.want {
			t.Errorf("Test for %d: Wanted: \n%vBut got:\n%v",test.input,test.want,output)
		}
	}
}
