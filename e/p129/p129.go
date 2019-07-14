package main

import (
	"fmt"
)

// This function computes A(n) and returns it for its argument n.
//
// A number consisting entirely of ones is called a repunit. 
// We shall define R(k) to be a repunit of length k; for example, R(6) = 111111.
// Given that n is a positive integer and GCD(n, 10) = 1, it can be shown that there always exists a value, k, 
// for which R(k) is divisible by n, and let A(n) be the least such value of k; for example, A(7) = 6 and A(41) = 5.
// The least value of n for which A(n) first exceeds ten is 17.
//
func A(n int) int {
	if n%2 == 0 || n%5 ==0 {
		return 0
	}
	k:=1
	i:=1
	for k%n != 0 {
		i++
		k=(10*k+1)%n
		// fmt.Printf("A: %d %d\n",k,i)
	}
	return i
}

// Computes the least value of N for which A(N) > limit l
func SmallestN(limit int) int {
	var current int
	// start just above limit with the first odd number 
	if limit%2 == 0 {
		current=limit+1
	} else {
		current=limit+2
	}
	for i:=A(current); i<=limit; {
		// fmt.Println(i,current)
		current+=2
		i=A(current)
		// fmt.Println(i,current)

	}
	return current
}

func main() {
	fmt.Printf("A(7) = %d\n",A(7))
	fmt.Printf("A(41) = %d\n",A(41))
	fmt.Printf("A(17) = %d\n",A(17))
	fmt.Printf("SmallestN(10) = %d\n",SmallestN(10))
	fmt.Printf("SmallestN(1000000) = %d\n",SmallestN(1000000))

}