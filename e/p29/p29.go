package main

import (
	"fmt"
)

func main() {

	powers := map[int][]int{ //special cases where the first number is a power of something else
		4:   {2, 2}, // 4 is 2 to the power of 2
		8:   {2, 3}, // ... etc
		9:   {3, 2},
		16:  {2, 4},
		25:  {5, 2},
		27:  {3, 3},
		32:  {2, 5},
		36:  {6, 2},
		49:  {7, 2},
		64:  {2, 6},
		81:  {3, 4}, // 81 is 3 to the power of 4
		100: {10, 2},
	}

	seen := make(map[[2]int]bool)

	for i := 2; i <= 100; i++ {

		p, ok := powers[i]
		if !ok { //simple case, add all of the powers in
			for j := 2; j <= 100; j++ {
				seen[[2]int{i, j}] = true
			}
			continue
		}

		for j := int(100/p[1]) + 1; j <= 100; j++ { //some powers would have already been added, add just those that could be bigger
			seen[[2]int{p[0], j * p[1]}] = true // here for example instead of adding 64 to the power of 17, we add 2 to the power of 6*17
		}
	}

	// uncomment to see all of the distinct powers 
	// fmt.Println(seen)
	fmt.Println(len(seen))

}
