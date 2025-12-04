package queues

import "fmt"

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