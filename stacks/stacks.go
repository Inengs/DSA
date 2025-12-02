package stacks

import (
	"errors"
)

const maxSize = 5 // maximum stack size

// ====================== Array-based Stack ======================

type ArrayStack struct {
	items [maxSize]int
	top   int // index of the next free position
}

// Push adds a value to the stack
func (s *ArrayStack) Push(value int) error {
	if s.top >= maxSize {
		return errors.New("stack overflow")
	}
	s.items[s.top] = value
	s.top++
	return nil
}

// Pop removes and returns the top value
func (s *ArrayStack) Pop() (int, error) {
	if s.top == 0 {
		return 0, errors.New("stack underflow")
	}
	s.top--
	return s.items[s.top], nil
}

// Top returns the top value without removing it
func (s *ArrayStack) Top() (int, error) {
	if s.top == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[s.top-1], nil
}

// Count returns the number of elements in the stack
func (s *ArrayStack) Count() int {
	return s.top
}

// ====================== Linked-list Stack ======================

// Node represents each element in the linked list stack
type Node struct {
	value int
	next  *Node
}

type LinkedStack struct {
	top   *Node // points to the top node
	count int   // number of elements
}

// Push adds a value to the top of the stack
func (s *LinkedStack) Push(value int) {
	newNode := &Node{
		value: value,
		next:  s.top,
	}
	s.top = newNode
	s.count++
}

// Pop removes and returns the top value
func (s *LinkedStack) Pop() (int, error) {
	if s.top == nil {
		return 0, errors.New("stack underflow")
	}
	value := s.top.value
	s.top = s.top.next
	s.count--
	return value, nil
}

// Top returns the top value without removing it
func (s *LinkedStack) Top() (int, error) {
	if s.top == nil {
		return 0, errors.New("stack is empty")
	}
	return s.top.value, nil
}

// Count returns the number of elements in the stack
func (s *LinkedStack) Count() int {
	return s.count
}
