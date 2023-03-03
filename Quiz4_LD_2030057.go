package main

import "fmt"

func reverseString(str *string) {
	// I found this way to make a string mutable
	// ussing byte slice to allow modification

	s := []byte(*str)
	length := len(s)
	// Using swapping characters
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
	*str = string(s)
}

func main() {

	var myString string = "Dale Duro"
	fmt.Println("Original string:", myString)
	reverseString(&myString)
	fmt.Println("Reversed string:", myString)
}
