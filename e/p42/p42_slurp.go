package main

import (
	"fmt"
	"math"
	"os"
	"io/ioutil"
	"strings"
)


func is_triangle_number(n int) bool {
	sr := math.Sqrt((float64(n*2)))
	if int(sr)*(int(sr)+1) == n*2 {
		return true
	} else {
		return false
	}
}

func main() {
	// fmt.Printf("Is triangle 45? %v\n",is_triangle_number(45))
	
	const filename = "/Users/top/Downloads/p042_words.txt"
	f, err :=os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	slurp, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	words:=strings.Split(string(slurp),",")

	fmt.Println(words)


	triangles :=0

	for _, word := range words{
		sum := 0
		for _, r:= range word {
			if r>='A' && r<='Z' {
				sum += int(r-'A')+1
			}
		}
		if is_triangle_number(sum) {
			triangles++
		}
	}

	fmt.Println("Got triangles: ",triangles)
}