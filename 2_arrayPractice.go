package main

import "fmt"

func main() {

	//Approach #1
	var arrayOfInt [5]int
	var arrayOfString [10]string
	var arrayOfFloat [3]float64
	var arrayOfBool [4]bool

	arrayOfInt[1] = 10
	arrayOfString[5] = "Luis El Grande"
	arrayOfFloat[0] = 10.98
	arrayOfBool[2] = true

	fmt.Println(arrayOfInt[1], ",", arrayOfString[5], ",", arrayOfFloat[0], ",", arrayOfBool[2])

	//Approach #2

	var arrayOfInt2 = [5]int{1, 12, 92, 36, 15}
	fmt.Println(arrayOfInt2)

	//Approach #3
	arrayOfFloat2 := [4]float64{1.2, -3.14, 0.0, 45.5}
	fmt.Println(arrayOfFloat2)
	//Approach #3
	//int[] array = new int[10]

	arrayOfString2 := [...]string{"Luis", "The", "Great"}
	fmt.Println(arrayOfString2)

	for i := 0; i < len(arrayOfString2); i++ {
		fmt.Println(arrayOfString2[i])
	}
}
