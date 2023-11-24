package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Linkedlist struct {
	headnode *Node
	length   int
}

func (li *Linkedlist) prepend(n *Node) {
	temp := li.headnode
	li.headnode = n
	n.next = temp
	li.length++
}

func (li *Linkedlist) Printall() {

	current := li.headnode

	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (li *Linkedlist) delete(v int) {
	prev := li.headnode
	if li.headnode.value == v {
		li.headnode.next = li.headnode.next.next
		return
	}

	for prev.next.value != v {
		fmt.Println("hellothere")
		if prev.next == nil {
			fmt.Println("nf")
			return
		}
		prev = prev.next
	}
	prev.next = prev.next.next
	li.Printall()
}

func main() {

	li := Linkedlist{}
	n1 := Node{value: 5}
	n2 := Node{value: 6}
	n3 := Node{value: 7}
	n4 := Node{value: 8}
	li.prepend(&n1)
	li.prepend(&n2)
	li.prepend(&n3)
	li.prepend(&n4)
	li.delete(6)
	li.Printall()
}
