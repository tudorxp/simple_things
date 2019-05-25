package main

// A simple program that prints out a Pascal triangle
// (which is computed using a singly linked list)

import (
	"fmt"
	"flag"
)

type ll_elem struct {
	value int
	next *ll_elem
}

func main() {

	var iterations int
	flag.IntVar(&iterations,"i",10,"number of iterations, default 10")
	flag.Parse()

	fmt.Println("Iterations: ",iterations)

	list := &ll_elem {1, nil}

	for i:=1; i<=iterations; i++ {
		// print
		for x:=list; x!=nil; x=x.next {
			fmt.Print(x.value," ")
		}
		fmt.Println()

		// new head
		new_head := new(ll_elem)
		new_head.value=1
		new_head.next=list
		list=new_head

		// increment
		for x:=list.next; x!=nil; x=x.next {
			if x.next!=nil {
				x.value+=x.next.value
			}
		}
	}
}