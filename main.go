package main

import (
	"fmt"
	tree "main/node"
	"main/node/queues/valq"
)

const to = 9

func main() {
	t := tree.GetTree(1)
	for i := 2; i < to+1; i++ {
		t.Add(i)
	}
	t.Val.LeftToRight()
	for i := 1; i < 5; i += 2 {
		fmt.Println("deleting:", i)
		if i == 5 {
			valq.Debug = true
		}
		t.Delete(i)
	}
	fmt.Println(t)
	t.Val.LeftToRight()
	//t.Delete(2)
}
