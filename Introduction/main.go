package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Linkedlist struct {
	headnode *Node
}

func (li *Linkedlist) prepend(n *Node) {
	temp := li.headnode
	li.headnode = n
	n.next = temp
}

func (li Linkedlist) Printall() {
	toprint := li.headnode
	for toprint != nil {
		fmt.Println(toprint.value)
		toprint = toprint.next
	}
}

func (li Linkedlist) Search(v int) {
	current := li.headnode
	var pos int = 0
	for current.next != nil {
		if current.value != v {
			current = current.next
			pos++
		} else {
			fmt.Println(pos)
			return
		}

	}
}

func (li *Linkedlist) Delete(v int) {
	prevtodel := li.headnode
	if prevtodel.value == v {
		li.headnode = prevtodel.next
		return
	}
	for prevtodel.next.value != v {
		if prevtodel.next == nil {
			return
		}
		prevtodel = prevtodel.next
	}
	prevtodel.next = prevtodel.next.next
}

func (li *Linkedlist) reftolast() *Node {
	current := li.headnode
	if current == nil {
		fmt.Println("null")
		return nil
	}
	for current.next != nil {
		current = current.next
	}
	return current
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

func (li *Linkedlist) rotate() {
	current := li.headnode
	for current.next != nil {
		fmt.Println("currentnext", current.next)
		temp := current.value
		current.next.value = temp
		fmt.Println("currentnext", current.next)
		current = current.next
		fmt.Println("curernt", current)
	}
	li.headnode.value = current.value

}

func main() {

	n1 := Node{value: 3}
	n2 := Node{value: 5}
	n3 := Node{value: 7}
	list := Linkedlist{}
	list.prepend(&n1)
	list.prepend(&n2)
	list.prepend(&n3)
	list.Printall()
	list.rotate()
	list.Printall()

}
