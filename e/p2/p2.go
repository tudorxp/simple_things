package main

import (
	"fmt"
)


// Finds the sum of all even Fibonacci numbers <= limit
func fibonacci_evens_sum(limit int) int {
	sum:=0
	for f1, f2 := 1, 1; f1+f2<=limit; f1, f2 = f2, f1+f2 {
		if (f1+f2)%2==0 {
			sum+=f1+f2
		}
	}
	return sum
}

func main() {
	fmt.Println("Sum (10):",fibonacci_evens_sum(4000000))
}