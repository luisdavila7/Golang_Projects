package main

import "fmt"

type sampleStructure struct {
	variable1 string
	variable2 string
	variable3 string
	variable4 int
	variable5 int
	variable6 float64
}

func return1(myStruct *sampleStructure) (string, string, float64) {
	return myStruct.variable1, myStruct.variable2, myStruct.variable6
}

func return2(myStruct *sampleStructure) (string, int, int) {
	return myStruct.variable3, myStruct.variable4, myStruct.variable5
}

func return3(myStruct *sampleStructure) float64 {
	return myStruct.variable6
}

func main() {

	myStruct := sampleStructure{"Mr", "Luis", "Davila", 35, 95, 120000.00}

	var1, var2, var3 := return1(&myStruct)
	var4, var5, var6 := return2(&myStruct)
	var7 := return3(&myStruct)

	fmt.Println(var1, var2, var3)
	fmt.Println(var4, var5, var6)
	fmt.Println(var7)

}
