package main

import (
	"fmt"
)

type sieve struct {
	// 3126*64 = at max 200064 bit entries
	// each entry is 1 if the corresponding number has been found to NOT be a prime number
	data [3126]uint64
}

// Get returns the value of the n-th bit
func (s sieve) Get(n uint64) uint64 {
	byte := uint64(n / 64)
	return (s.data[byte] >> (n % 64)) & 1
}

// Set modifies the value of the n-th bit to x
func (s *sieve) Set(n uint64, x uint64) {
	byte := uint64(n / 64)
	switch x {
	case 1:
		s.data[byte] |= 1 << (n % 64)
	default:
		s.data[byte] &= ^(1 << (n % 64))
	}

}

func main() {
	var s sieve

	primes := 0
	var i uint64

	for i = 2; primes <= 10000; i++ {
		if s.Get(i) == 0 { //found prime
			primes++
			for j := i * 2; j <= 200000; j += i {
				s.Set(j, 1)
			}
		}
	}

	fmt.Println(i - 1)

}
