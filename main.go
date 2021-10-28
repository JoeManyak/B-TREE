package main

import (
	"fmt"
	tree "main/node"
	"math/rand"
)

func main() {
	t := tree.GetTree(1)
	//rand.Seed(time.Now().UnixNano())
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(3)
	for i := 0; i < 10000; i++ {
		t.Add(rand.Intn(10000))
	}
	fmt.Println(t)
	/*t := tree.GetTree(5)
	t.Add(3)
	t.Add(7)
	t.Add(2)
	t.Add(4)
	t.Add(3)
	//valq.Debug = true
	t.Add(2)
	//fmt.Println("-")
	t.Add(8)
	t.Add(9)*/
	t.Val.LeftToRight()
	/*t.Add(1)
	t.Add(2)
	t.Add(3)*/
	/*rand.Seed(time.Now().UnixNano())
	for i:=0;i<100;i++{
		t.Add(rand.Intn(1000))
	}*/

	//fmt.Println(t)
	//	t.Add(9)
	/*	gen.InsertVal(1,true)
			gen.InsertVal(8,true)
		a,b,c:=gen.GetMedian()
		fmt.Println(gen,a,b,c)*/
}
