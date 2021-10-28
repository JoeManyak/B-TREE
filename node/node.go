package tree

import (
	"fmt"
	"main/node/queues/valq"
)

type Tree struct {
	Val *valq.ValQueue
}

func GetTree(value int) Tree {
	gen := valq.GenValQ(value)
	newTree := Tree{gen}
	return newTree
}

func (t *Tree) Add(val int) {
	fmt.Println("adding", val)
	_, link := t.Val.InsertVal(val, true)
	link.Subdivide()
	t.Val = t.Val.GetRoot()
}
