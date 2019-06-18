package main

import (
	"fmt"
	"strconv"
)

func is_palindrome(n int) bool {
	s := strconv.Itoa(n)
	b := []rune(s)

	for i:=0; i<=len(b)/2; i++ {
		if b[i]!=b[len(b)-1-i] {
			return false
		}
	}
	return true
}

func sum_squares(start, end int) int {
	sum:=0
	for i:=start; i<=end; i++ {
		sum+=i*i
	}
	return sum
}

func main () {
	sums:=make(map[int]bool)

	const upperlimit=1e8

	var s int

	for s=1; s*s + (s+1)*(s+1) <= upperlimit; s++ {
		for j:=s+1; ; j++ {
			ss := sum_squares(s,j)
			if ss>upperlimit {
				break
			}
			if is_palindrome(ss) {
				sums[ss]=true
			}
		}
	}

	sum:=0
	for x, _:= range sums {
		sum+=x
	}
	// fmt.Println(sums)
	fmt.Println(sum)

}
