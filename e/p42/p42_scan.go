package main

import (
	"fmt"
	"math"
	"os"
	"bufio"
	"io"
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

	words := []string{}

	r:=bufio.NewReader(f)

	for s, err := r.ReadString(',') ; err == nil || err == io.EOF; s, err = r.ReadString(',') {
		fmt.Println(s)
		s:=strings.Trim(s,"\",")
		fmt.Println(s)

		words=append(words,s)
		if err==io.EOF {
			break
		}
	}

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