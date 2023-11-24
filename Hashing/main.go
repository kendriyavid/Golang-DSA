package main

import "fmt"

const hl = 10

type Node struct {
	next  *Node
	value string
}

type Linkedlist struct {
	headnode *Node
	length   int
}

type Hashtable struct {
	Hasharray [hl]*Linkedlist
}

func (L *Linkedlist) insertNode(v string) {
	temp := L.headnode
	newnode := &Node{value: v}
	L.headnode = newnode
	newnode.next = temp
	L.length++
}

func (L *Linkedlist) deletenode(v string) *Linkedlist {
	prev := L.headnode
	if L.headnode.value == v {
		L.headnode = L.headnode.next
		L.length--
		return L
	}
	for prev.next.value != v {
		prev = prev.next
	}
	prev.next = prev.next.next
	L.length--
	return L
}

func (H *Hashtable) Hashinsert(v string) {
	hash := Hash(v)
	H.Hasharray[hash].insertNode(v)
}

func (H *Hashtable) PrintIndex(v int) {
	if H.Hasharray[v].length == 0 {
		fmt.Println("nil")
	}
	current := H.Hasharray[v].headnode
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}

}

func (H *Hashtable) Delete(v string) {
	hash := Hash(v)
	laddr := H.Hasharray[hash].deletenode(v)
	H.Hasharray[hash] = laddr
}

func (H *Hashtable) init() {
	for i := 0; i < hl; i++ {
		H.Hasharray[i] = &Linkedlist{}
	}
}

func Hash(v string) int {
	len := len(v)
	sum := 0
	for i := 0; i < len; i++ {
		sum = int(v[i]) + sum
	}
	hash := sum % hl
	return hash
}

func main() {
	hashar := &Hashtable{}
	hashar.init()
	hashar.Hashinsert("harsh")
	fmt.Println(hashar)
	hashar.PrintIndex(4)
	hashar.Delete("harsh")
	fmt.Println("hithere")
	hashar.PrintIndex(4)

}
