package main

import "fmt"

func main() {

	var ptr *int

	fmt.Println(ptr)

	var myVal = 8.91
	var ptr2 = &myVal

	fmt.Println("Value of myVal =", myVal)
	fmt.Println("Adress of myVal = ", &myVal)
	fmt.Println("Value where ptr2 is refering to = ", ptr2)
	fmt.Println("==========================================")
	fmt.Printf("Address of ptr2 = %#x \n", ptr2)

	fmt.Println("==========================================")

	ptr3 := new(float64)
	*ptr3 = 4.3

	fmt.Printf("ptr3 address = %#x, ptr3 value = %f\n", ptr3, *ptr3)

	fmt.Println("==========================================")

	var myVal2 = 101
	var ptr4 = &myVal2
	var ptr5 = &ptr4

	fmt.Println("myVal2", myVal2)
	fmt.Println("Address of myVal2 = ", &myVal2)

	fmt.Println("ptr4 = ", ptr4)
	fmt.Println("Address of ptr4 = ", &ptr4)

	fmt.Println("ptr5 = ", ptr5)
	fmt.Println("vale of this pointer = ", **ptr5)

	fmt.Println("====================================")

	var myVal3 = 101
	var ptr6 = &myVal3
	fmt.Println(ptr6)

	if *ptr5 == ptr6 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// var ptr7 = ptr6 + 1

}
