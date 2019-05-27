package main

import (
	"fmt"
)

func die_if_error(err error){
	if err!=nil {
		panic(err)
	}
}

func find_sum(n,s int, a []int) (int, int) {
	i:=0
	j:=0
	sum:=a[i]

	for {

		// fmt.Printf("%d %d %d\n",i,j,sum)

		switch {
		
		case sum==s:
			// Found it
			return i,j

		case sum<s:
			// Shifting end of window
			j++
			if j<n {
				sum+=a[j]
			} else {
				return -1,-1
			}

		case sum>s:
			// Shifting beginning of window
			sum-=a[i]
			i++
			if i<n {
				continue
			} else {
				return -1,-1
			}

		}

	}
}

func main() {

	var n,s int
	var a []int
	var t int
	var output []string

	// Get number of test cases
	_,err := fmt.Scan(&t)
	output=make([]string,t)

	die_if_error(err)

	for i:=0;i<t;i++ {
		// Read each test case
		_,err:=fmt.Scan(&n,&s)
		die_if_error(err)

		a=make([]int,n)
		for j:=0;j<n;j++ {
			_,err:=fmt.Scan(&a[j])
			die_if_error(err)
		}
		// Execute
		b,e := find_sum(n,s,a)

		if b==-1 {
			output=append(output,fmt.Sprintln(b))
		} else {
			output=append(output,fmt.Sprintln(b+1,e+1))
		}
	}

	for _,s := range output {
		fmt.Print(s)
	}

}