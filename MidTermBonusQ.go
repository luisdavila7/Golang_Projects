package main

import "fmt"

type myQueue struct {
	array []int
	left  int
	right int
}

func (q *myQueue) Enqueue(value int) {
	q.array = append(q.array, value)
	q.right = len(q.array) - 1
	if q.left == -1 {
		q.left = 0
	}
}

func (q *myQueue) Dequeue() int {
	if q.left == -1 {
		return -1
	}
	value := q.array[q.left]
	if q.left == q.right {
		q.left = -1
		q.right = -1
	} else {
		q.left++
	}
	return value
}

func main() {
	q := &myQueue{
		array: make([]int, 0),
		left:  -1,
		right: -1,
	}

	q.Enqueue(5)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	fmt.Println("Queue after Enqueue: ", q.array)
	fmt.Println("Dequeued item: ", q.Dequeue())
	fmt.Println("Queue after Dequeue: ", q.array)
	fmt.Println("Dequeued item: ", q.Dequeue())
}
