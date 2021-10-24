package tree

import "main/node/queues/valq"

type Tree struct {
	Val *valq.ValQueue
}

func GetTree(value int) Tree {
	gen := valq.GenValQ(5)
	newTree := Tree{gen}
	return newTree
}

func (t *Tree) Add(val int) {
	t.Val.InsertVal(val, true)
	t.Val = t.Val.GetRoot()
}
