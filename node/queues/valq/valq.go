package valq

import "fmt"

const max = 2

var Debug = false

type ValQueue struct {
	FirstVal  *ValNode
	FirstLink *LinkNode
	Parent    *ValQueue
	Count     int
	isLeaf    bool
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
	fmt.Println(l)
	newNode := LinkNode{
		Link: link,
		Next: l.Next,
	}
	l.Next = &newNode
}

func GenValQ(value int) *ValQueue {
	return &ValQueue{
		FirstVal: &ValNode{
			Val:  value,
			Next: nil,
		},
		isLeaf: true,
		Count:  1,
	}
}

func (v *ValQueue) InsertLink(id int, value int, link *ValQueue) {
	cur := v.FirstLink
	if v.FirstVal.Val > value {
		newNode := ValNode{
			Val:  value,
			Next: v.FirstVal,
		}
		v.FirstVal = &newNode
		return
	}
	for i := 0; i < id; i++ {
		cur = cur.Next
	}
	if cur.Link != nil {
		if cur.Link.FirstVal.Val > value {
			cur.Append(link)
		}
	}
	cur.Next.Append(link)
}

/*func (v *ValQueue) Raise() {
	median, left, right := v.GetMedian()

}*/

func (v *ValQueue) GetRoot() *ValQueue {
	cur := v
	for cur.Parent != nil {
		cur = cur.Parent
	}
	return cur
}

func (v *ValQueue) Subdivide() {
	if v.Count > max {
		fmt.Println("Subdividing... , count:", v.Count)
		median, left, right := v.GetMedian()
		if v.Parent == nil {
			v.Parent = &ValQueue{
				FirstVal: median,
				FirstLink: &LinkNode{
					Link: &ValQueue{
						FirstVal:  left,
						FirstLink: nil,
						Count:     max / 2,
						isLeaf:    true,
					},
					Next: &LinkNode{
						Link: &ValQueue{
							FirstVal:  right,
							FirstLink: nil,
							Count:     max / 2,
							isLeaf:    true,
						},
						Next: nil,
					},
				},
				Parent: nil,
				Count:  1,
				isLeaf: false,
			}
			v.Parent.FirstLink.Link.Parent = v.Parent
			v.Parent.FirstLink.Next.Link.Parent = v.Parent
		} else {
			id := v.Parent.InsertVal(median.Val, false)
			v.Parent.InsertLink(id, median.Val, v)
		}
	} else {
		fmt.Println("No need to subdivide, count:", v.Count)
	}
}

func (v *ValQueue) GetMedian() (*ValNode, *ValNode, *ValNode) {
	cur := v.FirstVal
	leftQVal := v.FirstVal
	for i := 0; i < v.Count/2-1; i++ {
		cur = cur.Next
	}
	median := cur.Next
	cur.Next = nil
	rightQVal := median.Next
	median.Next = nil
	return median, leftQVal, rightQVal
}

/*func (v *ValQueue) InsertLink() {

}*/

func (v *ValQueue) InsertVal(value int, isGoingDown bool) int {
	if v.isLeaf || !isGoingDown {
		v.Count++
		i := 0
		cur := v.FirstVal
		if cur.Val > value {
			newNode := ValNode{
				Val:  value,
				Next: v.FirstVal,
			}
			v.FirstVal = &newNode
			v.Subdivide()
			return i
		}
		for cur.Next != nil {
			if cur.Next.Val > value {
				cur.Next = &ValNode{
					Val:  value,
					Next: cur.Next,
				}
				v.Subdivide()
				return i
			}
			i++
			cur = cur.Next
		}

		cur.Next = &ValNode{
			Val:  value,
			Next: nil,
		}
		i++
		v.Subdivide()
		return i
	} else {
		curVal := v.FirstVal
		curLink := v.FirstLink
		for curVal.Next != nil {
			if curVal.Val > value {
				return curLink.Link.InsertVal(value, isGoingDown)
			}
			curVal = curVal.Next
			curLink = curLink.Next
		}
		if curVal.Val > value {
			return curLink.Link.InsertVal(value, isGoingDown)
		}
		curLink = curLink.Next
		return curLink.Link.InsertVal(value, isGoingDown)
	}
}
