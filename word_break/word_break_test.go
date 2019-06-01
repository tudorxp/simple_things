package main

import (
	"testing"
	// "fmt"
	"strings"
)

func Test_word_break(t *testing.T) {

	dictionary := map[string]bool{
		"i" :true, 
		"like":true,
		"sam": true,
		"sung":true, 
		"samsung":true,
		"mobile":true, 
		"ice":true,
		"cream":true, 
		"icecream":true, 
		"man":true, 
		"go":true, 
		"mango":true,
	}

	tests := []struct {
		input string
		possible bool
	}{
		struct { input string; possible bool;}{ 
			"ilikesamsungandmango",
			false,
		},
		struct { input string; possible bool}{ 
			"iiiiiiii",
			true,
		},
		struct { input string; possible bool}{ 
			"ilikelikeimangoiii",
			true,
		},
		struct { input string; possible bool}{ 
			"",
			true,
		},
	}

	for _,test:=range tests {
		p,s := word_break([]rune(test.input),dictionary)
		if p!=test.possible {
			t.Errorf("string %s, want=%v but got %v",test.input,test.possible,p)
			continue
		}
		if !p { continue }
		// check output string matches input
		if strings.Join(strings.Split(s," "),"") != test.input {
			t.Errorf("string %s, output string got: %s.",test.input,s)
		}
		// check all words are in dictionary
		for _,w := range strings.Split(s," ") {
			if len(w)!=0 && !dictionary[w] {
				t.Errorf("word not in dictionary: %s",w)				
			}
		}
	}

}