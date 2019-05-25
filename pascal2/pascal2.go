package main

import (
	"log"
	"flag"
)

func main() {
	iterations:=flag.Int("i",10,"number of iterations; default 10")
	flag.Parse()

	s:= make([]int,0,*iterations) 
	log.Printf("iterations: %d, type of s: %T",*iterations,s)

	for i:=0;i<*iterations;i++ {
		s=append(s,1)
		for j:=len(s)-2;j>=1;j-- {
			s[j]+=s[j-1]
		}
		log.Print(s)
	}

}