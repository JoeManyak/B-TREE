package main

import (
	"fmt"
	tree "main/node"
)

func main() {
	t := tree.GetTree(1)
	for i := 66; i > 2; i++ {
		t.Add(i)
	}
	t.Delete(2)
	fmt.Println(t)
}
