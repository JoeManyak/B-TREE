package main

import (
	tree "main/node"
	"math/rand"
	"time"
)

const to = 10000

func main() {
	//for i:= 0;i<100;i++ {
	tree.Deleted = 0
	tree.Added = 0
	//fmt.Println("start with seed:",int64(i))
	rand.Seed(time.Now().UnixNano())
	t := tree.GetTree(1)
	//counter := 0
	for i := 0; i < to; i++ {
		t.Add(rand.Intn(to))
		if tree.Deleted+t.Val.CountAll(0)-1 != tree.Added {
			//fmt.Println(tree.Deleted, t.Val.CountAll(0)-1, tree.Added)
			panic("wha1")
		}
	}

	//fmt.Println("switch=>",tree.Deleted, t.Val.CountAll(0)-1, tree.Added)
	/*for i := 0; i < to; i++ {
		t.Delete(rand.Intn(to))
		fmt.Println(tree.Deleted, t.Val.CountAll(0)-1, tree.Added)
		if tree.Deleted+t.Val.CountAll(0)-1 != tree.Added {
			panic("wha2")
		}
	}
	t.Val.LeftToRight()
	fmt.Println(tree.Deleted, t.Val.CountAll(0), tree.Added+1)
	CompCount := make([]int,15,15)
	for i:=0;i<15;i++ {
		queues.Comparisons = 0
		searching := rand.Intn(to)
		fmt.Println("search for:",searching)
		t.Search(searching)
		CompCount[i] = queues.Comparisons
	}
	fmt.Println(CompCount)
	sum:=0
	for _,v:=range CompCount{
		sum+=v
	}
	fmt.Println(sum)*/
	//panic("s")
}
