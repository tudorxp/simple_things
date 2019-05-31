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

// word_break takes arguments: a string as a slice of runes, and a dictionary
// it returns true, space separated words if the input string can be broken into words from the dictionary
// or false, empty string if not
func word_break(s []rune, dictionary map[string]bool) (bool, string) {

	// a[i].possible is true if for a specific index i, s[0:i+1] is word breakable
	// a[i].word is the (last) word that made it true
	var a []struct{
		possible bool
		word string
	}
	a=make([]struct {possible bool; word string}, len(s))

	// l is a map which stores all values of i where a[i]==true
	l:=make(map[int]bool)
	l[-1]=true // because null substring is true

	for i:=0; i<len(s); i++ {
		if a[i].possible != true {
			for j, _ := range l { // check substrings from prefixes which matched
				if dictionary[string(s[j+1:i+1])] {
					a[i].possible=true
					a[i].word=string(s[j+1:i+1])
					l[i]=true
					// fmt.Printf("Found match %d:%d: %s\n",j+1,i+1,string(s[j+1:i+1]))
					break
				}
			}
		}
	}

	if len(s)==0 {
		// fmt.Printf("Word break possible: empty string\n")
		return true, ""
	}
	if a[len(s)-1].possible==true {
		// fmt.Printf("Word break possible\n")
		// compose string
		words:=[]string{}
		for i:=len(s)-1;i>=0; {
			words=append(words,a[i].word)
			i-=len(a[i].word)
		}
		str:=""
		for i:=len(words)-1;i>=0;i-- { 
			str+=words[i]+" "
		}
		// fmt.Println(str)
		return true, str
	}
	return false, ""

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

	possible, str := word_break(s,dictionary)
	if possible {
		fmt.Println("Word break possible: ",str)
	} else {
		fmt.Println("Word break NOT possible")
	}

}