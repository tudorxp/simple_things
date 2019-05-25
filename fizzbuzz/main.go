package main

import (
	"fmt"
)

func main() {

	res := ""
	for i := 1; i <= 100; i++ {
		r := ""
		if i%3 == 0 {
			r += "fizz"
		}
		if i%5 == 0 {
			r += "buzz"
		}
		if r!="" {
			res += r+" "
			continue
		}
		res += fmt.Sprintf("%d ",i)
	}

	fmt.Println(res)

}
