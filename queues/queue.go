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


// implementation of linear queue using data structure
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