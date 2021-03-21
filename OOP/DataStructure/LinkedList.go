package DataStructure

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Root *Node
}

func (L *LinkedList) AddFirst(val interface{}) {
	curnode := &Node{val, L.Root}
	L.Root = curnode
}

func (L *LinkedList) AddLast(val interface{}) {
	if L.Root == nil {
		L.AddFirst(val)
	} else {
		curnode := L.Root
		for curnode.Next != nil {
			curnode = curnode.Next
		}
		curnode.Next = &Node{val, nil}
	}
}

func (L *LinkedList) PrintList() {
	curnode := L.Root
	for curnode.Next != nil {
		fmt.Print(curnode.Value)
		curnode = curnode.Next
	}
}
