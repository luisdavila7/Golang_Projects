package main

import "fmt"

func reverseString(str *string) {
	s := []byte(*str)
	length := len(s)
	// Using swapping characters
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
	*str = string(s)
}

func main() {
	var myString string
	fmt.Printf("Please insert a word to reverse: ")
	fmt.Scan(&myString)
	fmt.Println("Original string:", myString)
	reverseString(&myString)
	fmt.Println("Reversed string:", myString)
}
