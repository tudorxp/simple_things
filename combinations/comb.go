package main

import (
	"fmt"
)

func backtrack(n,m int, depth int, a []int, seen []bool) {
	switch depth {
	case m+1:
		fmt.Println(a)
	default:
		start:=0
		if depth>1 {
			start=a[depth-2]
		}
		for i:=start;i<n;i++ {
			if !seen[i] {
				a[depth-1]=i+1
				seen[i]=true
				backtrack(n,m,depth+1,a,seen)
				seen[i]=false
				a[depth-1]=0
			}
		}
	}
}

func main() {

	var n,m int
	fmt.Print("Enter n and m: ")
	_, err := fmt.Scanf("%d %d",&n,&m)
	if err != nil {
		panic(err)
	}

	a := make([]int, m)
	seen := make([]bool, n)

	backtrack(n,m,1,a,seen)
}