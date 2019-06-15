package main

import (
	"fmt"
)

// The first approach that comes to mind for counting the number of set bits is also the most readable;
// In the absence of performance requirements, the readable code has been included as the default solution.
//
// population_count(n) loops through all of the significant bits of n, checking them.
func population_count(n uint) (count uint) {
	for i := n; i > 0; { // start from n and continue until zero
		count += i & 1 // if the least significant bit is set then add it to the count
		i >>= 1 // bit shift right, discarding the bit we've just checked
	}
	return
}


func main() {
	tests := []uint{ 5, 0, 8, 15, 19 }
	for _, i := range tests {
		fmt.Printf("population_count(%d) = %d\n", i, population_count(i))
	}
}

// Main program ends here.






// Alternative approaches

// Another approach (Kernighan's, I believe) is to use the bit operation n & (n-1); on a non zero n, this will set to 0 the least significant bit set to 1
// Consequently, this method only needs <<count>> loops to complete, where count is the number of set bits.
func population_count_kernighan(n uint) (count uint) {
	for i := n; i > 0; {
		count++
		i = i & (i - 1)
	}
	return
}


// Yet another approach is to maintain a lookup table of already precomputed bit counts, which could be a 4-bit table (16 entries), 8-bit table (256 entries), etc.
// Here we are using a 4-bit table.
// NB: Go's standard library code in math/bits uses an 8-bit table for <=32 bit OnesCount functions.
var popcount_table16 = [...]uint { 0,1,1,2,1,2,2,3,1,2,2,3,2,3,3,4 }
func population_count_table(n uint) (count uint) {
	for i:=n; i>0; {
		count += popcount_table16[i & 0xf]
		i>>=4;
	}
	return
}

// Here is a constant time approach. It is by far the least readable, and requires a reference to an explanation: https://stackoverflow.com/a/15979139
// NB: Go's standard library code in math/bits uses a similar approach for 64-bit OnesCount.
func population_count_divideconquer(n uint) (count uint) {
	i := uint32(n) - ((uint32(n) >> 1) & 0x55555555)
    i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
    return uint((((i + (i >> 4)) & 0x0f0f0f0f) * 0x01010101) >> 24)
}