package valq

import "fmt"

const max = 2

var Debug = false

type ValQueue struct {
	FirstVal  *ValNode
	FirstLink *LinkNode
	Parent    *ValQueue
	Count     int
	IsLeaf    bool
}

type ValNode struct {
	Val  int
	Next *ValNode
}

type LinkNode struct {
	Link *ValQueue
	Next *LinkNode
}

func (l *LinkNode) Append(link *ValQueue) {
	newNode := LinkNode{
		Link: link,
		Next: l.Next,
	}
	l.Next = &newNode
}

func (v *ValQueue) LeftToRight() {
	cur := v.FirstLink
	val := v.FirstVal
	if cur != nil {
		for cur.Next != nil {
			cur.Link.LeftToRight()
			fmt.Println(val.Val)
			cur = cur.Next
			val = val.Next
		}
		cur.Link.LeftToRight()
	} else {
		for val != nil {
			fmt.Println(val.Val)
			val = val.Next
		}
	}
}

func GenValQ(value int) *ValQueue {
	return &ValQueue{
		FirstVal: &ValNode{
			Val:  value,
			Next: nil,
		},
		IsLeaf: true,
		Count:  1,
	}
}

func (v *ValQueue) InsertLink(id int, value int, link *ValQueue) {
	cur := v.FirstLink
	for i := 0; i < id; i++ {
		cur = cur.Next
	}
	cur.Append(link)
	/*cur := v.FirstLink
	if v.FirstLink.Link.FirstVal.Val >= value {
		newNode := LinkNode{
			Link: link,
			Next: v.FirstLink,
		}
		v.FirstLink = &newNode
		return
	}
	for i := 0; i < id; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		if cur.Next.Link.FirstVal.Val >= value {
			cur.Append(link)
			return
		}
	}
	cur.Append(link)*/
}

func (v *ValQueue) GetRoot() *ValQueue {
	cur := v
	for cur.Parent != nil {
		cur = cur.Parent
	}
	return cur
}

func (v *ValQueue) Subdivide() {
	if v.Count > max {
		//fmt.Println("Subdividing... , count:", v.Count)
		median, left, right, leftLink, rightLink := v.GetMedian()
		if v.Parent == nil {
			v.Parent = &ValQueue{
				FirstVal: median,
				FirstLink: &LinkNode{
					Link: &ValQueue{
						FirstVal:  left,
						FirstLink: leftLink,
						Count:     max / 2,
						IsLeaf:    leftLink == nil,
					},
					Next: &LinkNode{
						Link: &ValQueue{
							FirstVal:  right,
							FirstLink: rightLink,
							Count:     max / 2,
							IsLeaf:    rightLink == nil,
						},
						Next: nil,
					},
				},
				Parent: nil,
				Count:  1,
				IsLeaf: false,
			}
			v.Count = 1
			//parenting
			//parenting
			//parenting
			//parenting
			//parenting
			//parenting
			//parenting
			//cur := v.Parent.FirstLink.Next.Link
			/*if cur.FirstLink != nil {
				cur.FirstLink.Link.Parent = cur
				cur.FirstLink.Next.Link.Parent = cur
			}*/
			/*v.Parent.FirstLink.Link.Parent = v.Parent
			v.Parent.FirstLink.Next.Link.Parent = v.Parent*/

			v.Parent.ReParent()
			if Debug {
				t := v.GetRoot()
				fmt.Println(t)
			}
			curRepair := v.Parent.FirstLink
			if !v.IsLeaf {
				for curRepair != nil {
					curRepair.Link.ReParent()
					curRepair = curRepair.Next
				}
			}
			/*if !v.IsLeaf {
				v.Parent.FirstLink.Link.ReParent()
				v.Parent.FirstLink.Next.Link.ReParent()
			}*/
			if Debug {
				t := v.GetRoot()
				fmt.Println(t)
			}
		} else {
			id, _ := v.Parent.InsertVal(median.Val, false)
			v.Parent.IsLeaf = false
			v.Count = 1
			v.Parent.InsertLink(id, median.Val, &ValQueue{
				FirstVal:  right,
				FirstLink: rightLink,
				Parent:    v.Parent,
				Count:     max / 2,
				IsLeaf:    rightLink == nil,
			})

			v.Parent.ReParent()
			curRepair := v.Parent.FirstLink
			if !v.IsLeaf {
				for curRepair != nil {
					curRepair.Link.ReParent()
					curRepair = curRepair.Next
				}
			}
			v.Parent.Subdivide()
		}
	} else {
		//fmt.Println("No need to subdivide, count:", v.Count)
	}
}

func (v *ValQueue) ReParent() {
	cur := v.FirstLink
	for cur != nil {
		cur.Link.Parent = v
		cur = cur.Next
	}
}

func (v *ValQueue) GetMedian() (*ValNode, *ValNode, *ValNode, *LinkNode, *LinkNode) {
	cur := v.FirstVal
	leftQVal := v.FirstVal
	for i := 0; i < v.Count/2-1; i++ {
		cur = cur.Next
	}
	median := cur.Next
	cur.Next = nil
	rightQVal := median.Next
	median.Next = nil
	v.IsLeaf = v.FirstLink == nil
	if v.IsLeaf {
		return median, leftQVal, rightQVal, nil, nil
	} else {
		leftLink := v.FirstLink
		curL := v.FirstLink
		for i := 0; i < v.Count/2; i++ {
			curL = curL.Next
		}

		newCur := curL.Next
		curL.Next = nil
		rightLink := newCur
		return median, leftQVal, rightQVal, leftLink, rightLink
	}
}

/*func (v *ValQueue) InsertLink() {

}*/

func (v *ValQueue) InsertVal(value int, isGoingDown bool) (int, *ValQueue) {
	if v.IsLeaf || !isGoingDown {
		v.Count++
		i := 0
		cur := v.FirstVal
		if cur.Val >= value {
			newNode := ValNode{
				Val:  value,
				Next: v.FirstVal,
			}
			v.FirstVal = &newNode
			return i, v
		}
		i++
		for cur.Next != nil {
			if cur.Next.Val >= value {
				cur.Next = &ValNode{
					Val:  value,
					Next: cur.Next,
				}
				return i, v
			}
			i++
			cur = cur.Next
		}

		cur.Next = &ValNode{
			Val:  value,
			Next: nil,
		}
		//v.Subdivide()
		return i, v
	} else {
		curVal := v.FirstVal
		curLink := v.FirstLink
		for curVal.Next != nil {
			if curVal.Val >= value {
				return curLink.Link.InsertVal(value, isGoingDown)
			}
			curVal = curVal.Next
			curLink = curLink.Next
		}
		if curVal.Val >= value {
			return curLink.Link.InsertVal(value, isGoingDown)
		}
		curLink = curLink.Next
		return curLink.Link.InsertVal(value, isGoingDown)
	}
}

//:=:=:=:=:=:=:=:=:=: DELETIONS :=:=:=:=:=:=:=:=:=:\\

func (v *ValQueue) Search(value int) (*ValQueue, int) {
	cur := v.FirstVal
	curL := v.FirstLink
	var i int
	for i = 0; cur != nil; i++ {
		if !v.IsLeaf {
			if value <= cur.Val {
				val, count := curL.Link.Search(value)
				if val != nil {
					return val, count
				}
			}
		}
		if cur.Val == value {
			return v, i
		}
		cur = cur.Next
		if !v.IsLeaf {
			curL = curL.Next
		}
	}
	if !v.IsLeaf {
		val, count := curL.Link.Search(value)
		if val != nil {
			return val, count
		}
	}
	return nil, 0
}

func (v *ValQueue) DeleteLink(id int) {
	if id == 0 {
		v.FirstLink = v.FirstLink.Next
	} else {
		cur := v.FirstLink
		for i := 0; i < id-1; i++ {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}
}

func (v *ValQueue) DeleteVal(id int) {
	if id == 0 {
		v.FirstVal = v.FirstVal.Next
	} else {
		cur := v.FirstVal
		for i := 0; i < id-1; i++ {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}
	v.Count--
}

func (v *ValQueue) DeleteAndBalance(id int) {
	v.DeleteVal(id)
	if v.IsLeaf {
		if v.FirstVal != nil {
			return
		}
		selfId := v.GetSelfLinkId()
		left, right := v.GetNeighbours()
		if left != nil {
			if left.Count > 1 {
				fromParent := v.Parent.Get(selfId - 1)
				biggest := left.GetBiggest()
				v.InsertVal(fromParent.Val, false)
				v.Parent.InsertVal(biggest, false)
				return
			}
		}
		if right != nil {
			if right.Count > 1 {
				fromParent := v.Parent.Get(selfId)
				smallest := right.GetSmallest()
				v.InsertVal(fromParent.Val, false)
				v.Parent.InsertVal(smallest, false)
				return
			}
		}
		if left != nil {
			fromParent := v.Parent.Get(selfId - 1)
			left.InsertVal(fromParent.Val, true)
			v.Parent.DeleteLink(selfId)
			return
		}
		if right != nil {
			fromParent := v.Parent.Get(selfId)
			right.InsertVal(fromParent.Val, true)
			v.Parent.DeleteLink(selfId)
			return
		}
	} else {
		if v.FirstVal != nil {
			cur := v.FirstLink
			for i := 0; i < id; i++ {
				cur = cur.Next
			}
			if cur.Link.Count > 1 {
				val := cur.Link.GetBiggest()
				v.InsertVal(val, false)
				return
			}
			cur = cur.Next
			if cur.Link.Count > 1 {
				val := cur.Link.GetSmallest()
				v.InsertVal(val, false)
				return
			}
			value := cur.Link.Get(0)
			v.InsertVal(value.Val, true)
			v.DeleteLink(id + 1)
		} else {
			selfId := v.GetSelfLinkId()
			left, right := v.GetNeighbours()
			if left != nil {
				if left.Count > 1 {
					fromParent := v.Parent.Get(selfId - 1)
					biggest := left.GetBiggest()
					v.InsertVal(fromParent.Val, false)
					v.Parent.InsertVal(biggest, false)
					return
				}
			}
			if right != nil {
				if right.Count > 1 {
					fromParent := v.Parent.Get(selfId)
					smallest := right.GetSmallest()
					v.InsertVal(fromParent.Val, false)
					v.Parent.InsertVal(smallest, false)
					return
				}
			}
		}
	}
}

func (v *ValQueue) GetSelfLinkId() int {
	if v.Parent != nil {
		cur := v.Parent.FirstLink
		for i := 0; i < v.Parent.Count+1; i++ {
			if cur.Link == v {
				fmt.Println("")
				return i
			}
			cur = cur.Next
		}
	}
	return 0
}

func (v *ValQueue) Get(id int) *ValNode {
	cur := v.FirstVal
	for i := 0; i < id; i++ {
		cur = cur.Next
	}
	v.DeleteVal(id)
	return cur
}

func (v *ValQueue) GetSmallest() int {
	smallest := v.FirstVal.Val
	v.DeleteVal(0)
	return smallest
}

func (v *ValQueue) GetBiggest() int {
	cur := v.FirstVal
	for cur.Next != nil {
		cur = cur.Next
	}
	biggest := cur.Val
	v.DeleteVal(v.Count - 1)
	return biggest
}

func (v *ValQueue) GetNeighbours() (*ValQueue, *ValQueue) {
	if v.Parent == nil {
		return nil, nil
	}
	var i int
	cur := v.Parent.FirstLink
	for i = 0; cur != nil; i++ {
		if cur.Link == v {
			break
		}
		cur = cur.Next
	}
	var left, right *ValQueue
	cur = v.Parent.FirstLink
	for j := 0; cur != nil; j++ {
		if j == i-1 {
			left = cur.Link
		}
		if j == i+1 {
			right = cur.Link
			return left, right
		}
		cur = cur.Next
	}
	return left, right
}
