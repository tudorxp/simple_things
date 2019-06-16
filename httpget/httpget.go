package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	const url = "http://lpeak.com"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	for k, v := range resp.Header {
		fmt.Printf("Got header: %s, values: %v\n",k,v)
	}
	fmt.Println("Response code:",resp.StatusCode)

	content, err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		panic(err)
	}
	fmt.Println("Content: ")
	fmt.Println(string(content))

}
