package main

import "fmt"

func main() {

	var myMap1 map[string]string
	var myMap2 map[string]int
	var myMap3 map[int]float64

	fmt.Println(myMap1, myMap2, myMap3)

	var myMap4 = make(map[string]float64)
	fmt.Println(myMap4)

	var myMap5 = map[string]string{
		"Day0": "Poutine",
		"Day1": "Baneja Paisa",
		"Day2": "Kebab",
		"Day3": "Taco",
	}

	fmt.Println(myMap5["Day0"])

	var studentGPA = map[string]float64{
		"Luis":   4.5,
		"Aldo":   3.9,
		"Behnam": 2.91,
		"Prince": 3.11,
	}

	fmt.Println(studentGPA)
	delete(studentGPA, "Aldo")
	fmt.Println(studentGPA)

	for key, gpa := range studentGPA {
		if gpa == 3.11 {
			fmt.Println(key)
		}
	}
}
