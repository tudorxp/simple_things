package main

import (
	"fmt"
	"strings"
)

// pascal computes a pascal triangle of depth n
// it returns a string containing the output for a given depth like so:
// 1\n
// 1 1\n
// 1 2 1\n
func pascal(n int) string {
	var b strings.Builder
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, 1)
		for j := i - 1; j > 0; j-- {
			s[j] = s[j] + s[j-1]
		}
		fmt.Fprintf(&b, "%v\n", s)
	}
	return b.String()
}

func main() {
	var iterations int
	fmt.Printf("Enter number of iterations: ")
	_, err := fmt.Scanf("%d",&iterations)
	if err != nil {
		panic(err)
	}

	fmt.Print(pascal(iterations))
}
