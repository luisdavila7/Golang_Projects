package main

import "fmt"

const MAX_SIZE = 100

type Stack struct {
	arr []interface{}
	top int
}

func createStack() *Stack {
	stack := &Stack{arr: make([]interface{}, MAX_SIZE), top: -1}
	for i := 0; i < MAX_SIZE; i++ {
		stack.arr[i] = nil
	}
	return stack
}

func push(stack *Stack, value interface{}) {
	if stack.top == MAX_SIZE-1 {
		fmt.Println("Error: Stack is full!")
		return
	}
	stack.top++
	stack.arr[stack.top] = value
}

func pop(stack *Stack) interface{} {
	if stack.top == -1 {
		fmt.Println("Error: Stack is empty!")
		return nil
	}
	poppedValue := stack.arr[stack.top]
	stack.arr[stack.top] = nil
	stack.top--
	return poppedValue
}

func main() {
	stack := createStack()

	push(stack, 10)
	push(stack, "Dale")
	push(stack, "Duro")
	push(stack, "SVP")

	fmt.Println(pop(stack))
	fmt.Println(pop(stack))
	fmt.Println(pop(stack))
	fmt.Println(stack)
}
