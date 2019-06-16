package main

import (
	"fmt"
)

// we compute the sum of all numbers <1000 that either divide by 3 or 5
func main() {
	sum := 0
	for i:=3; i<1000; i+=3 {
		sum += i
	}
	for i:=5; i<1000; i+=5 {
		if i%3 != 0 { // already added in previous loop
			sum += i
		}
	}
	fmt.Println("Sum is: ",sum)
}