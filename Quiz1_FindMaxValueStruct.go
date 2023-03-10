package main

import "fmt"

type numbers struct {
	number int
}

func findthe3MaxValues(array [10]numbers) {

	var max1 = 0
	var max2 = 0
	var max3 = 0
	var ptr1 = &max1
	var ptr2 = &max2
	var ptr3 = &max3

	fmt.Println("Array is => ", array)

	for i := 0; i < len(array); i++ {
		if array[i].number > *ptr1 {
			max3 = max2
			max2 = max1
			max1 = array[i].number
		} else if array[i].number > *ptr2 {
			if max1 == array[i].number {
			} else {
				max3 = max2
				max2 = array[i].number
			}
		} else if array[i].number > *ptr3 {
			if *ptr1 == array[i].number || *ptr2 == array[i].number {
			} else {
				max3 = array[i].number
			}
		}
	}

	fmt.Printf("The maximun 3 numbers are => %d, %d, %d", max1, max2, max3)
}

func main() {

	var arrayNumbers [10]numbers

	arrayNumbers[1] = numbers{1}
	arrayNumbers[2] = numbers{2}
	arrayNumbers[3] = numbers{3}
	arrayNumbers[4] = numbers{7}
	arrayNumbers[5] = numbers{8}
	arrayNumbers[6] = numbers{9}
	arrayNumbers[7] = numbers{10}
	arrayNumbers[8] = numbers{12}
	arrayNumbers[9] = numbers{90}
	arrayNumbers[0] = numbers{5}

	findthe3MaxValues(arrayNumbers)

}
