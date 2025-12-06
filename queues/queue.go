package queues

import "fmt"

// implementation for linear queue using array
const maxSize int = 5

var queue [maxSize]int

var front int = -1
var rear int = -1

func enqueue(x int) {
	if rear == maxSize-1 {
		fmt.Printf("Overflow")
	} else if front == -1 && rear == -1 {
		front = 0
		rear = 0
		queue[rear] = x
	} else {
		rear++
		queue[rear] = x 
	}
}

func dequeue() {
	if front == -1 && rear == -1 {
		fmt.Println("Underflow")  
	} else if front == rear{
		front = -1
		rear = -1
	} else {
		fmt.Printf("%d\n", queue[front])
		front++
	}
}

func display() {
	if front == -1 && rear == -1 {
		fmt.Println("queue is empty")
	} else {
		for i := 0; i<rear+1; i++ {
			fmt.Printf("%d\n", queue[i])
		}
	}
}

func peek() {
	if front == -1 && rear == -1 {
		fmt.Println("queue is empty")
	} else { 
		fmt.Printf("%d\n", queue[front])	
	}
}


// implementation of linear queue using singly linked-list
type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	front *Node 
	rear *Node
}

func NewQueue() *Queue {
	return &Queue{} // CREATES AN EMPTY QUEUE
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{Value: value, Next: nil}

	if q.front == nil && q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.Next = newNode
		q.rear = newNode
	}
}

func (q *Queue) Display() {
	var temp *Node
	if q.front == nil && q.rear == nil {
		fmt.Println("queue is empty")
	} else {
		temp = q.front
		for temp != nil{
			fmt.Printf("%d ", temp.Value)
			temp = temp.Next
		}
		fmt.Println()
	}
}

func (q *Queue) Dequeue() {
	if q.front == nil && q.rear == nil {
		fmt.Println("queue is empty")
		return
	}

	q.front = q.front.Next

	if q.front == nil {
		q.rear = nil
	}
}

func (q *Queue) Peek() {
	if q.front == nil && q.rear == nil {
		fmt.Println("queue is empty")
	} else {
		fmt.Printf("%d", q.front.Value)
	}
}

// using stacks
type Stack []int

type QueueForStack struct {
	inStack  Stack
	outStack Stack
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1 // or error handling
	}
	n := len(*s)
	val := (*s)[n-1]
	*s = (*s)[:n-1]
	return val
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (q *QueueForStack) Enqueue(v int) {
	q.inStack.Push(v)
}


func (q *QueueForStack) Dequeue() int {
	// If outStack is empty, move everything from inStack
	if q.outStack.IsEmpty() {
		for !q.inStack.IsEmpty() {
			val := q.inStack.Pop()
			q.outStack.Push(val)
		}
	}

	// Now outStack has the oldest element at its top
	return q.outStack.Pop()
}

func (q *QueueForStack) Front() int {
	if q.outStack.IsEmpty() {
		for !q.inStack.IsEmpty() {
			val := q.inStack.Pop()
			q.outStack.Push(val)
		}
	}

	// Top of outStack is the front of queue
	n := len(q.outStack)
	return q.outStack[n-1]
}
