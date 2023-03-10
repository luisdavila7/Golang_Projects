package main

import "fmt"

type myEmployees struct {
	employeId           *int
	empName             *string
	empLastName         *string
	empJob              *string
	empSupervisor       *string
	empSalary           *float64
	empEmergencyContact *string
	empInsurance        *string
	empBank             *string
	isActiveEmp         *bool
}

var arrayofmyEmployees [1]myEmployees

func main() {

	ptr1 := new(int)
	*ptr1 = 100

	ptr2 := new(string)
	*ptr2 = "Luis"

	ptr3 := new(string)
	*ptr3 = "Davila"

	ptr4 := new(string)
	*ptr4 = "Developer"

	ptr5 := new(string)
	*ptr5 = "Max"

	ptr6 := new(float64)
	*ptr6 = 120000

	ptr7 := new(string)
	*ptr7 = "Laurent"

	ptr8 := new(string)
	*ptr8 = "Dejardins"

	ptr9 := new(string)
	*ptr9 = "BMO"

	ptr10 := new(bool)
	*ptr10 = true

	arrayofmyEmployees[0] = myEmployees{ptr1, ptr2, ptr3, ptr4, ptr5, ptr6, ptr7, ptr8, ptr9, ptr10}

	fmt.Println("Emp Id: ", *arrayofmyEmployees[0].employeId)
	fmt.Println("Emp Name: ", *arrayofmyEmployees[0].empName)
	fmt.Println("Emp LastName: ", *arrayofmyEmployees[0].empLastName)
	fmt.Println("Emp Job: ", *arrayofmyEmployees[0].empJob)
	fmt.Println("Emp Supervisor: ", *arrayofmyEmployees[0].empSupervisor)
	fmt.Println("Emp Salary: ", *arrayofmyEmployees[0].empSalary)
	fmt.Println("Emp Contact: ", *arrayofmyEmployees[0].empEmergencyContact)
	fmt.Println("Emp Insurance: ", *arrayofmyEmployees[0].empInsurance)
	fmt.Println("Emp Bank: ", *arrayofmyEmployees[0].empBank)
	fmt.Println("Emp Is Active: ", *arrayofmyEmployees[0].isActiveEmp)

}
