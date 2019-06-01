package main

import (
	"os"
	"flag"
	"encoding/json"
	"fmt"
	"bufio"
)

func die_if_error(err error) {
	if err!=nil {
		panic(err)
	}
}

// word_break takes arguments: a string as a slice of runes, and a dictionary; plus index for the backtracking, and a pointer to a slice of words already found
// it prints out all possible word break sequences
func word_break(s []rune, dictionary map[string]bool, index int, words *[]string) {
	switch {
	case index > len(s): //overshot
		return
	case index == len(s): //found something, print solution
		fmt.Println(*words)
		return
	default: //try every word from current index, if it matches add to words and recurse
		for w,_ := range dictionary {
			if string(s[index:index+len(w)])==w { //matches
				*words=append(*words,w)
				word_break(s,dictionary,index+len(w),words)
				*words=(*words)[:len(*words)-1]
			}
		}
	}
}

func main () {


	// Create dictionary
	dictionary_filename:=flag.String("dictionary","dict.json","dictionary (as JSON array) filename")
	flag.Parse()

	dictionary:=make(map[string]bool)

	f, err := os.Open(*dictionary_filename)
	die_if_error(err)
	defer f.Close()

	var dict_entries []string
	dec := json.NewDecoder(f)
	err = dec.Decode(&dict_entries)
	die_if_error(err)

	for i:=range dict_entries {
		dictionary[dict_entries[i]]=true
	}

	// fmt.Println(dictionary)

	// Get input string
	fmt.Print("Enter input string: ")
	var s []rune
	sc:=bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		panic("error reading input string")
	}
	// Convert to slice of runes
	s=[]rune(sc.Text())

	// fmt.Println("Got: ",string(s))

	words:=[]string{}
	word_break(s,dictionary,0,&words)
	
}