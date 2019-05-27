package main

import "testing"

func Test_find_sum(t *testing.T) {
	var n,s int
	var a []int
	var b,e int
	var rb,re int

	n,s = 5, 12
	a = []int{1,2,3,7,5}
	b,e = 1,3
	if rb,re=find_sum(n,s,a); rb!=b || re!=e {
		t.Errorf("failed, got=%d %d, want=%d %d",rb,re,b,e)
	}


	n,s = 10, 15
	a = []int{1,2,3,4,5,6,7,8,9,10}
	b,e = 0,4
	if rb,re=find_sum(n,s,a); rb!=b || re!=e {
		t.Errorf("failed, got=%d %d, want=%d %d",rb,re,b,e)
	}


	n,s = 10, 150
	a = []int{1,2,3,4,5,6,7,8,9,10}
	b,e = -1,-1
	if rb,re=find_sum(n,s,a); rb!=b || re!=e {
		t.Errorf("failed, got=%d %d, want=%d %d",rb,re,b,e)
	}


}
