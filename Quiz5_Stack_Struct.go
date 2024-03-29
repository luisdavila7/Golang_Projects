package main

import "fmt"

const MAX_SIZE = 100

type Stack struct {
	arr []int
	top int
}

func createStack() *Stack {
	stack := &Stack{arr: make([]int, MAX_SIZE), top: -1}
	return stack
}

func queueLeft(stack *Stack, value int) {
	if stack.top == MAX_SIZE-1 {
		fmt.Println("Error: Stack is full!")
		return
	}
	stack.top++
	stack.arr[stack.top] = value
}

func queueRight(stack *Stack, value int) {
	if stack.top == MAX_SIZE-1 {
		fmt.Println("Error: Stack is full!")
		return
	}
	stack.top++
	stack.arr[stack.top] = value
}

func deQueue(stack *Stack) int {
	if stack.top == -1 {
		fmt.Println("Error: Stack is empty!")
		return -1
	}
	poppedValue := stack.arr[stack.top]
	stack.top--
	return poppedValue
}

func main() {
	stack := createStack()

	queueLeft(stack, 10)
	queueRight(stack, 20)
	queueRight(stack, 30)
	queueLeft(stack, 40)

	fmt.Println(deQueue(stack))
	fmt.Println(deQueue(stack))
	fmt.Println(deQueue(stack))
	fmt.Println(stack)
}
