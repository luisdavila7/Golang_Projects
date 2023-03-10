package main

import (
	"fmt"
	"reflect"
)

func main() {

	var firstName string = "Luis"
	var lastName string = "Davila"
	var gpa float64 = 4.2
	var age int = 26
	//var isLaSalleMember bool = true

	var (
		SIN       int     = 123456
		lastName2 string  = "The great"
		isCoolGuy bool    = true
		salary    float64 = 120000
	)

	fmt.Println("First Name, Last Name, gpa and age are: ", firstName, lastName, gpa, age)

	fmt.Printf("SIN is %d and lastname is %s and %t and %f", SIN, lastName2, isCoolGuy, salary)

	fullname := "Luis El Grande"
	CurrentYear, income, isProgramer := 2023, 800.23, true
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("%s %d %f %t", fullname, CurrentYear, income, isProgramer)
	fmt.Println("")
	fmt.Println(reflect.TypeOf(fullname))
}
