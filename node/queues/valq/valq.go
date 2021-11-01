package valq

import "fmt"

const max = 10

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

func (v *ValQueue) InsertLink(id int, link *ValQueue) {
	cur := v.FirstLink
	if id == 0 {
		v.FirstLink = &LinkNode{
			Link: link,
			Next: cur,
		}
	} else {
		for i := 0; i < id-1; i++ {
			cur = cur.Next
		}
		cur.Append(link)
	}
	v.ReParent()
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
		if cur.Parent.Count > 0 {
			cur = cur.Parent
		} else {
			return cur
		}
	}

	for cur.Count == 0 {
		if cur.FirstLink != nil {
			cur = cur.FirstLink.Link
		} else {
			return nil
		}
	}
	return cur
}

func (v *ValQueue) Subdivide() {
	if v.Count > max {
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
		} else {
			id, _ := v.Parent.InsertVal(median.Val, false)
			v.Parent.IsLeaf = false //тупо
			v.Count = 1
			v.Parent.InsertLink(id+1, &ValQueue{
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

func (v *ValQueue) InsertVal(value int, isGoingDown bool) (int, *ValQueue) {
	if v.IsLeaf || !isGoingDown {
		v.Count++
		i := 0
		if v.FirstVal == nil {
			v.FirstVal = &ValNode{
				Val:  value,
				Next: nil,
			}
			return 0, v
		}
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

func (v *ValQueue) Balance(id int) {
	if v.IsLeaf {
		if v.FirstVal != nil {
			return
		}
		selfId := v.GetSelfLinkId()
		left, right := v.GetNeighbours()
		if left != nil {
			if left.Count > 1 {
				fromParent := v.Parent.Get(selfId - 1) //-1 косяк
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
			v.Parent.Balance(0)
			return
		}
		if right != nil {
			fromParent := v.Parent.Get(selfId)
			right.InsertVal(fromParent.Val, true)
			v.Parent.DeleteLink(selfId)
			v.Parent.Balance(0)
			return
		}
	} else {
		cur := v.FirstLink
		linkCount := 0
		for cur != nil {
			cur = cur.Next
			linkCount++
		}
		if v.Count == 0 && linkCount == 1 && v.Parent == nil {
			v.FirstLink.Link.Parent = nil
			return
		}
		if v.Count > 0 && v.Count+1 == linkCount {
			return
		}
		//Фіксим по стандарту, якщо далі,
		//то лівої дитини крайнє праве а далі
		//балансуємо в ноль)00)0
		if v.FirstLink.Next != nil {
			from, num := v.GetLeftBiggest(id)
			v.InsertVal(num, false)
			from.Balance(0)
			return
		}
		left, right := v.GetNeighbours()
		selfId := v.GetSelfLinkId()
		if left != nil {
			if left.Count > 1 {
				fromParent := v.Parent.Get(selfId - 1)
				biggest := left.GetBiggest()
				v.InsertVal(fromParent.Val, false)
				cur := left.FirstLink
				for cur.Next.Next != nil {
					cur = cur.Next
				}
				needed := cur.Next
				cur.Next = nil
				v.InsertLink(0, needed.Link)
				v.ReParent()
				v.Parent.InsertVal(biggest, false)
				return
			}
		}
		if right != nil {
			if right.Count > 1 {
				fromParent := v.Parent.Get(selfId)
				smallest := right.GetSmallest()
				v.InsertVal(fromParent.Val, false)
				v.InsertLink(1, right.FirstLink.Link)
				v.ReParent()
				right.FirstLink = right.FirstLink.Next
				v.Parent.InsertVal(smallest, false)
				return
			}
		}
		if left != nil {
			fromParent := v.Parent.Get(selfId - 1)
			left.InsertVal(fromParent.Val, false)
			left.InsertLink(left.Count, v.FirstLink.Link) //danger +1
			left.ReParent()
			v.Parent.DeleteLink(1)
			v.Parent.Balance(0)
			return
		}
		if right != nil {
			fromParent := v.Parent.Get(selfId)
			right.InsertVal(fromParent.Val, false)
			right.InsertLink(0, v.FirstLink.Link)
			right.ReParent()
			v.Parent.DeleteLink(0)
			v.Parent.Balance(0)
			return
		}

	}
}

func (v *ValQueue) GetRightest() (*ValQueue, int) {
	if v.IsLeaf {
		return v, v.GetBiggest()
	}
	cur := v.FirstLink
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur.Link.GetRightest()
}

func (v *ValQueue) GetLeftBiggest(id int) (*ValQueue, int) {
	cur := v.FirstLink
	for i := 0; i < id; i++ {
		cur = cur.Next
	}
	return cur.Link.GetRightest()
}

func (v *ValQueue) GetSelfLinkId() int {
	if v.Parent != nil {
		cur := v.Parent.FirstLink
		for i := 0; i < v.Parent.Count+1; i++ {
			if cur.Link == v {
				return i
			}
			cur = cur.Next
		}
	}
	panic("not found self")
	return -1
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
