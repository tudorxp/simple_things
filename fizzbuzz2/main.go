package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) (result string) {

	result = ""
	for i:=1; i<=n; i++ {

		r:=""
		if i%3==0 {
			r+="fizz"
		}
		if i%5==0 {
			r+="buzz"
		}
		if r!="" {
			result += r+" "
			continue
		}		
		result += strconv.Itoa(i)+" "
	}
	return

}

func main() {

	var n int
	fmt.Print("How much should I count to? ")
	fmt.Scanf("%d",&n)

	fmt.Print(FizzBuzz(n))

}