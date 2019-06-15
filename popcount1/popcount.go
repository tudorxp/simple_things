package main

import (
	"fmt"
	"math/bits"
)

func population_count_1(n uint) (count uint) {
	for i := n; i > 0; i >>= 1 {
		count += i & 1
	}
	return
}

func population_count_2(n uint) (count uint) {
	for i := n; i > 0; {
		count++
		i = i & (i - 1)
	}
	return
}

func population_count_3(n uint) (count uint) {
	return uint(bits.OnesCount32(uint32(n)))
}

func population_count_4(n uint) (count uint) {
	return uint(bits.OnesCount64(uint64(n)))
}


func main() {
	tests := []uint{5, 0, 8, 15, 19}
	for _, i := range tests {
		fmt.Printf("population_count(%d) = %d\n", i, population_count_1(i))
	}
}
