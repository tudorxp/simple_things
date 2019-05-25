package main

import (
	"fmt"
)

func backtrack (n int, depth int, a *[]int) {
	switch depth {
		case n+1:
			fmt.Println(*a)
		default:
			for i:=0; i<n; i++ {
				if (*a)[i]==0 {
					(*a)[i]=depth
					backtrack(n,depth+1,a)
					(*a)[i]=0
				}
			}
	}
}

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scanf("%d",&n)
	a := make([]int,n)
	backtrack(n,1,&a)
}