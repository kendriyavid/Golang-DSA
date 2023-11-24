package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	value string
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

func (li *Linkedlist) Delete(v string) {
	prev := li.headnode
	if li.headnode.value == v {
		li.headnode = li.headnode.next
		return
	}

	for prev.next.value != v {
		if prev.next == nil {
			fmt.Println("notfound")
			return
		}
		prev = prev.next
	}
	prev.next = prev.next.next
}

func (li *Linkedlist) reftolast() *Node {
	current := li.headnode

	for current.next != nil {
		current = current.next
	}
	return current
}

func (li *Linkedlist) rotate(v int) {
	lenght := li.lenght
	rotate := v % lenght
	current := li.headnode
	for current.next != nil {
		current = current.next
	}
	current.next = li.headnode

	traverse := lenght - rotate
	traversecurr := li.headnode
	for i := 0; i < traverse; i++ {
		traversecurr = traversecurr.next
	}
	li.headnode = traversecurr.next
	traversecurr.next = nil

}

func (li *Linkedlist) DeleteMiddle() {
	var current int = li.lenght / 2
	prev := li.headnode
	for i := 0; i < current-1; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next

}

func (li *Linkedlist) DeleteMiddle2() {
	fast := li.headnode
	var prev *Node
	slow := li.headnode
	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}
	prev.next = prev.next.next
}

func search(N *Node, arr [4]*Node) bool {
	for i := 0; i < 4; i++ {
		if arr[i] == N {
			return true
		} else if arr[i] != N {
			continue
		}
	}
	return false
}

func (li *Linkedlist) detectloop() {
	var hasharray [4]*Node
	current := li.headnode
	i := 0
	for current != nil {
		bool := search(current, hasharray)
		if bool == false {
			hasharray[i] = current
			current = current.next
			i++
		} else if bool == true {
			fmt.Println("presnet")
			return
		}
	}
	fmt.Println("no loop")

}

func (li *Linkedlist) detectloop2() {
	slow := li.headnode
	fast := li.headnode
	for slow == fast {
		slow = slow.next
		fast = fast.next.next
		if fast == nil || fast.next == nil {
			fmt.Println("no loop")
			return

		}
	}
	fmt.Println("loop present")
}

func (li *Linkedlist) Deleteloop() { // FLOYD CYCLE FINDING ALGO (SLOW AND FAST STARTING)
	var hasharray [4]*Node
	current := li.headnode
	var prev *Node = nil
	i := 0
	for current.next != nil {

		bool := search(current, hasharray)
		if bool == false {
			hasharray[i] = current
			i++
		} else if bool != false {
			prev.next = nil
			return
		}
		prev = current
		current = current.next
	}

}

func (li *Linkedlist) reverse(cli *Linkedlist) *Linkedlist {
	var prev *Node = nil
	list := *cli
	current := list.headnode
	for current != nil {
		temp := current.next
		current.next = prev
		prev = current
		current = temp
	}
	list.headnode = prev
	return &list
}

func (li *Linkedlist) reverseknodes(lit *Linkedlist, k int) {
	current := lit.headnode
	i := 1
	for current != nil {
		if i%k != 0 {
			current = current.next
			i++
		} else if i%k == 0 {
			temp := current
			current.next = nil
			reverse()
		}
	}
}

func main() {

	li := Linkedlist{}
	n1 := Node{value: "harshdeep"}
	n2 := Node{value: "Neetusingh"}
	n5 := Node{value: "pankhuri"}
	n4 := Node{value: "pankhuri"}

	li.prepend(&n1)
	li.prepend(&n2)
	li.prepend(&n4)
	li.prepend(&n5)
	li.headnode.next.next.next.next = li.headnode.next
	li.Deleteloop()
	li.printall()
}
