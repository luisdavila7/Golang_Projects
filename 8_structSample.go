package main

import "fmt"

type Human struct {
	age         int
	name        string
	beingMember bool
	gpa         float64
}

func main() {
	var obj1 Human
	fmt.Println(obj1)

	obj1.name = "Luis"
	obj1.age = 20
	obj1.gpa = 4.5
	obj1.beingMember = true
	fmt.Println(obj1)

	fmt.Println("==================")

	obj2 := Human{24, "Fer", true, 4.2}
	fmt.Println(obj2)
	fmt.Println(obj2.name)

}
