package main

import "fmt"
import "math"

func secuenceNumber(n int) {
	var val int
	val = 1
	for i := 0; i < n; i++ {
		val = int(math.Pow(2, float64(i)))
		fmt.Printf("%d, ", val)
	}
}

func main() {

	var n int
	fmt.Printf("Please insert how many number for the secuence: ")
	fmt.Scan(&n)
	secuenceNumber(n)
}
