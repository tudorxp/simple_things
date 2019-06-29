package main

import (
	"fmt"
	"tudorxp/simple_things/tree/tree"
)

func main() {
	// t:=tree.New(5,tree.New(3,nil,tree.New(4,nil,nil)),tree.New(8,tree.New(6,nil,tree.New(7,nil,nil)),nil))
	t:=tree.New(5,nil,nil)
	t.Add(3)
	t.Add(4)
	t.Add(8)
	t.Add(6)
	t.Add(7)



	fmt.Println(t.Value())
	fmt.Println(t.Find(4))
	fmt.Println(t.Find(6))
	fmt.Println(t.PreOrderTraversal())
	fmt.Println(t.InOrderTraversal())
	fmt.Println(t.PostOrderTraversal())


}