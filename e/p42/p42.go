package main

import (
	"fmt"
	"math"
	"encoding/csv"
	"os"
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

	r :=csv.NewReader(f)

	words, err :=r.Read()
	if err != nil {
		panic(err)
	}

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