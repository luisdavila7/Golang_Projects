package main

import "fmt"

func returnPointer(ptr1 *int, ptr2 *int) *int {

	ptr3 := 100

	return &ptr3
}

func main() {

	var val1 = 0
	var val2 = 0
	var ptr1 = &val1
	var ptr2 = &val2

	var ptr3 = returnPointer(ptr1, ptr2)

	fmt.Println(*ptr3)
}
