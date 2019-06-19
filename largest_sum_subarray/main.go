package main

import (
	"errors"
	"fmt"
	"math"
)

func largest_sum(a []int) (int, error) {
	if len(a) == 0 {
		return 0, errors.New("empty array has no sum")
	}
	s := a[0]
	for i := 0; i < len(a); i++ {
		cs := 0

		for j := i; j < len(a); j++ {
			cs += a[j]
			if cs > s {
				s = cs
			}
		}

	}
	return s, nil
}

func largest_sum_v2(a []int) (int, error) {
	if len(a) == 0 {
		return 0, errors.New("empty array has no sum")
	}

	s := 0

	cs := 0
	have_zero := false
	biggest_negative := math.MinInt64

	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			have_zero = true
		}
		if a[i] < 0 && a[i] > biggest_negative {
			biggest_negative = a[i]
		}
		cs += a[i]
		if cs > s {
			s = cs
		}
		if cs < 0 {
			cs = 0
		}

	}

	switch {
	case s != 0:
		return s, nil
	case s == 0 && have_zero:
		return s, nil
	default:
		return biggest_negative, nil
	}

}

func main() {
	test := []int{-5,-3,2,3,1,3,-4,-4,4,6,-1}
	sum, err := largest_sum_v2(test)
	if err != nil {
		panic(err)
	}
	fmt.Println("Largest sum is:", sum)
}
