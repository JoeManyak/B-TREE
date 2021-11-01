package main

import (
	tree "main/node"
	"math/rand"
	"time"
)

const to = 100000

func main() {
	t := tree.GetTree(1)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < to; i++ {
		t.Add(rand.Intn(to))
	}
	for i := 0; i < to/10; i++ {
		t.Delete(rand.Intn(to))
	}
	t.Val.LeftToRight()
}
