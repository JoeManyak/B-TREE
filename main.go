package main

import (
	"fmt"
	tree "main/node"
	"main/node/queues/valq"
)

func main() {
	t := tree.GetTree(5)
	t.Add(3)
	t.Add(7)
	t.Add(2)
	valq.Debug = true
	t.Add(1)
	fmt.Println(t)
	/*	gen.InsertVal(1,true)
		gen.InsertVal(8,true)
		a,b,c:=gen.GetMedian()
		fmt.Println(gen,a,b,c)*/
}
