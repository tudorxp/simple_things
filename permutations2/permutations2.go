package main

import (
	"fmt"
)

func backtrack (a []int, depth int, n int, v map[int]bool) {
	switch depth {
	case n:
		fmt.Println(a)
	default:
		for i:=0; i<n; i++ {
			if !v[i] {
				a[depth]=i
				v[i]=true
				backtrack(a,depth+1,n,v)
				v[i]=false
			}
		}
	}
}

func main() {
	fmt.Printf("Enter n: ")
	var n int
	_, err := fmt.Scanf("%d",&n)
	if err!=nil { panic(err) }
	a:=make([]int,n)
	v:=map[int]bool{}
	backtrack(a,0,n,v)
}