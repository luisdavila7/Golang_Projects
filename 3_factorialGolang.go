package main

import "fmt"

func factorialGo(number int) int {

	var fac int = 1

	if number == 1 {
		fac = 1
	} else {
		for i := 1; i < number; i++ {
			fac = fac + (fac * i)
		}
	}

	return fac
}

func main() {

	var result = factorialGo(5)
	fmt.Println(result)
}
