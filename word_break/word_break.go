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

func word_break(s []rune, dictionary map[string]bool) {
	var a []struct{
		possible bool
		word string
	}
	a=make([]struct {possible bool; word string}, len(s))
	_=a
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

	word_break(s,dictionary)

}