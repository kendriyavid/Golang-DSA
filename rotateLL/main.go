package main

import (
	"fmt"
)

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

func (li *Linkedlist) rotate(v int) {

	lenght := li.length
	rotate := v % lenght
	current := li.headnode
	for current.next != nil {
		current = current.next
	}
	current.next = li.headnode

	traverse := lenght - rotate
	tcurr := li.headnode
	for i := 0; i < traverse; i++ {
		tcurr = tcurr.next
	}
	li.headnode = tcurr.next
	tcurr.next = nil

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
	li.rotate(3)
	li.Printall()

}
