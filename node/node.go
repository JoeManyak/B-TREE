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
	fmt.Println("adding:", val)
	_, link := t.Val.InsertVal(val, true)
	link.Subdivide()
	t.Val = t.Val.GetRoot()
}

func (t *Tree) Search(val int) *valq.ValQueue {
	v, _ := t.Val.Search(val)
	return v
}

func (t *Tree) Delete(val int) {
	search, id := t.Val.Search(val)
	if search != nil {
		search.DeleteVal(id)
		search.Balance(id)
	} else {
		fmt.Println("Not found :", val)
	}
	t.Val = t.Val.GetRoot()
}
