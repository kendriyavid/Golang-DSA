package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

type Linkedlist struct {
	headnode *Node
	lenght   int
}

func (li *Linkedlist) prepend(n *Node) {
	temp := li.headnode
	li.headnode = n
	n.next = temp
	li.lenght++
}

func (li *Linkedlist) printall() {
	current := li.headnode
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (li *Linkedlist) reverse() {
	current := li.headnode
	var prev *Node = nil
	for current != nil {
		temp := current.next
		current.next = prev
		prev = current
		current = temp
	}
	li.headnode = prev
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
	li.reverse()
	li.printall()

}
