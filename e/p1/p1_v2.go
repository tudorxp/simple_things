package main

import (
	"fmt"
)

// we compute the sum of all numbers <1000 that either divide by 3 or 5
func main() {
	sum := 0

	sum = 3*(333*334/2) // sum of all numbers that divide by 3
	sum += 5*(199*200/2) // sum of all numbers that divide by 5
	sum -= 15*(66*67/2) // sum of all numbers that divide by both 3 and 5

	fmt.Println("Sum is: ",sum)
}