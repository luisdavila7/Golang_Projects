package main

import "fmt"

type PointerArgs struct {
	ptr1 *int
	ptr2 *int
}

func ReturnPointer(args *PointerArgs) *int {
	ptr3 := 100
	return &ptr3
}

func main() {
	var val1, val2 int
	var ptr1 = &val1
	var ptr2 = &val2
	args := &PointerArgs{ptr1, ptr2}
	ptr3 := ReturnPointer(args)
	fmt.Println(*ptr3)
}
