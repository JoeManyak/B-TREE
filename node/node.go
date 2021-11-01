package tree

import (
	"fmt"
	"main/node/queues"
)

const printing = true

var Deleted = 0
var Added = 0

type Tree struct {
	Val *queues.ValQueue
}

func GetTree(value int) Tree {
	gen := queues.GenValQ(value)
	newTree := Tree{gen}
	return newTree
}

func (t *Tree) Add(val int) {
	Added++
	if printing {
		fmt.Println("adding:", val)
	}
	_, link := t.Val.InsertVal(val, true)
	link.Subdivide()
	t.Val = t.Val.GetRoot()
}

func (t *Tree) Search(val int) *queues.ValQueue {
	v, _ := t.Val.Search(val)
	return v
}

func (t *Tree) Delete(val int) {
	search, id := t.Val.Search(val)
	if search != nil {
		Deleted++
		if printing {
			fmt.Println("Deleting :", val)
		}
		search.DeleteVal(id)

		search.Balance(id)
	} else {
		if printing {
			fmt.Println("Not found :", val)
		}
	}
	t.Val = t.Val.GetRoot()
}
