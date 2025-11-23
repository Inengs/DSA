package doubly

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{} // returns a new linkedList with head set to nil
}

func (ll *LinkedList) InsertAtHead(value int) bool {
	newNode := &Node{Value: value, Next: ll.head, Prev: nil}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.head.Prev = newNode
		ll.head = newNode
	}

	ll.size++
	return true
}

func (ll *LinkedList) InsertAtSpecific(value int, position int) bool {
	if position == 0 {
		return ll.InsertAtHead(value) // Handle edge cases
	}

	newNode := &Node{Value: value, Next: nil, Prev: nil} // create a new node with the given data

	var current = ll.head //

	for i := 0; i < position-1; i++ { // loop/tranverse through the to position-1
		if current == nil {
			return false // invalid position
		}

		current = current.Next // Move to the next, keep on going
	}

	//Link the new node, now you have reach where you want to insert
	newNode.Next = current.Next
	newNode.Prev = current

	// if current's next exists, update its previous pointer
	if current.Next != nil {
		current.Next.Prev = newNode // let its previous pointer point to the newNode
	}

	//Update current' next to new Node
	current.Next = newNode
	ll.size++
	return true
}

// Helper function to get size
func (ll *LinkedList) GetSize() int {
	return ll.size
}

// Delete at specific position
func (ll *LinkedList) DeleteAtPosition(position int) bool {
	// Check if position is valid
	if position < 0 || position >= ll.size {
		return false
	}

	// Traverse to the node at given position
	current := ll.head
	for i := 0; i < position; i++ {
		current = current.Next
	}

	// Update links to remove current node
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev

	ll.size--
	return true
}

// Search by value (forward)
func (ll *LinkedList) Search(value int) bool {
	current := ll.head

	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}

	return false
}

// Traverse forward
func (ll *LinkedList) TraverseForward() {
	current := ll.head

	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}