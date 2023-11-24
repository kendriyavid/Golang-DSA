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

func main() {
	li := Linkedlist{}
	n1 := Node{value: "Ford"}
	n2 := Node{value: "ferrari"}
	n5 := Node{value: "pagani"}
	n4 := Node{value: "lexus"}

	li.prepend(&n1)
	li.prepend(&n2)
	li.prepend(&n4)
	li.prepend(&n5)
	li.headnode.next.next.next.next = li.headnode.next
	li.Deleteloop()
	li.printall()
}
