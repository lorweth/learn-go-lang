package main

import (
	structure "OOP/DataStructure"
)

func main() {
	l := structure.LinkedList{}
	l.AddFirst("A")
	l.AddFirst(1)
	l.AddFirst("C")
	l.AddLast(3)
	l.PrintList()
}
