package main

import (
	"fmt"
)

func bsort(a []int) {
	for i:=0; i<len(a); i++ {
		changed:=false
		for j:=0; j<len(a)-i-1; j++ {
			if a[j]>a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				changed=true
			}
		}
		if changed == false {
			break 
		}
		fmt.Println(a)		
	}
}

func main() {
	a := []int{2,3,11,5,1,3,9,3,58,3}
	fmt.Println(a)
	bsort(a)
	fmt.Println(a)
}