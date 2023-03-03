package main

import "fmt"

type Student struct {
	studentId     string
	fullname      string
	age           int
	gpa           float64
	laSalleMember bool
}

func main() {

	std1 := Student{"12345", "Luis", 35, 94, true}
	fmt.Println(std1)

	ptr := &std1
	fmt.Println("===========================")
	fmt.Println(ptr, &std1)
	fmt.Println(&ptr)
	fmt.Println((*ptr).studentId, (*ptr).fullname)

	//ptr.gpa = 4.1
	fmt.Println(ptr.gpa)

}
