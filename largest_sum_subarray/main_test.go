package main

import (
	"testing"
)

var r int

func Benchmark_largest_sum(b *testing.B) {
	for i:=0; i<b.N; i++ {
		res, _:=largest_sum([]int{-5,-3,2,3,1,3,-4,-4,4,6,-1})
		r=res
	}
}


func Benchmark_largest_sum_v2(b *testing.B) {
	for i:=0; i<b.N; i++ {
		res, _:=largest_sum_v2([]int{-5,-3,2,3,1,3,-4,-4,4,6,-1})
		r=res
	}
}


// [top@tmac t1]$ go test -bench .
// goos: darwin
// goarch: amd64
// pkg: tudorxp/t1
// Benchmark_largest_sum-8      	30000000	        58.3 ns/op
// Benchmark_largest_sum_v2-8   	100000000	        23.2 ns/op
// PASS
// ok  	tudorxp/t1	4.165s