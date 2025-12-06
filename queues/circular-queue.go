package queues

import "fmt"

// implementation using arrays
func enqueueCircular(x int) {
	if front == -1 && rear == -1 {
		front = 0
		rear = 0
		queue[rear] = x
	} else if (rear + 1) % maxSize == front{
		fmt.Println("Queue is full")
	} else {
		rear = (rear + 1) % maxSize
		queue[rear] = x
	} 
}

func dequeueCircular(x int) {
	if front == -1 && rear == -1{

	} else if front == rear {
		front = -1
		rear = -1
	} else {
		fmt.Printf("%d", queue[front])
		front = (front + 1) % maxSize
	}
}

func displayCircular() {
	if(front == -1 && rear == -1) {
		fmt.Println("queue is empty")
	} else {

		fmt.Printf("%d", )
		i := front
		for i != rear{
			fmt.Printf("%d", queue[i])
			i = (i+1)%maxSize
		}
	}
}

// implementation using linked-lists
func (q *Queue) EnqueueCircular(value int) {
	newNode := &Node{Value: value, Next: nil}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		q.rear.Next = q.front
	} else {
		q.rear.Next = newNode
		q.rear = newNode
		q.rear.Next = q.front 
	}
}

func(q *Queue) DequeueCircular() {
	if q.front == nil {
		return
	}

	// If only one element
	if q.front == q.rear {
		q.front = nil
		q.rear = nil
		return
	}

	// Move front forward
	q.front = q.front.Next
	q.rear.Next = q.front // maintain circle
}

func (q *Queue) DisplayCircular() {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return
	}

	temp := q.front
	for {
		fmt.Printf("%d ", temp.Value)
		temp = temp.Next
		if temp == q.front {
			break
		}
	}
	fmt.Println()
}
