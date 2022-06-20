package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	val  interface{}
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) append(val interface{}) {
	node := Node{
		val: val,
	}
	if list.head == nil {
		list.head = &node
	} else {
		ptr := list.head
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = &node
		node.prev = ptr
	}
}

func (list *LinkedList) printList() {
	ptr := list.head
	for ptr != nil {
		fmt.Printf("Node(%v)->", ptr.val)
		ptr = ptr.next
	}
	fmt.Println("nil")
}
