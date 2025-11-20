package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{} // returns a new linkedList with head set to nil
}

func (ll *LinkedList) Search(value int) bool {
	var current = ll.head

	for current != nil {
		if current.Value == value {
			return true
		} else {
			current = current.Next
		}
	}

	return false
}

func (ll *LinkedList) Delete(value int) bool {
	var current = ll.head
	var previous *Node = nil

	if ll.head == nil {
		return false
	}

	if ll.head.Value == value {
		ll.head = ll.head.Next
	}

	for current != nil {
		if current.Value == value {
			previous.Next = current.Next
			return true
		} else {
			previous = current
			current = current.Next
		}
	}

	return false
}

func (ll *LinkedList) InsertAtTail(value int) bool {
	newNode := &Node{Value: value, Next: nil}
	var current = ll.head

	if ll.head == nil {
		ll.head = newNode
		return true
	} else {
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		return true
	}
}

func (ll *LinkedList) Transverse() {
	var current = ll.head

	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (ll *LinkedList) ReverseTransverse(node *Node) { // call with ll.head when you want to use
	if node == nil {
		return
	}

	ll.ReverseTransverse(node.Next)
	fmt.Println(node.Value)
}


func (ll *LinkedList) InsertAtTop(value int) {
	newNode := &Node{Value: value, Next: ll.head}
	ll.head = newNode
}