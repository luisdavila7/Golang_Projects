package main

import (
	"fmt"
	"math"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Error: Division by zero!")
	}
	return a / b, nil
}

func modulo(a, b float64) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Error: Modulo by zero!")
	}
	return int(a) % int(b), nil
}

func logarithm(a float64) (float64, error) {
	if a <= 0 {
		return 0, fmt.Errorf("Error: Invalid input!")
	}
	return math.Log10(a), nil
}

func squareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("Error: Invalid input!")
	}
	return math.Sqrt(a), nil
}

func main() {
	var choice int
	var num1, num2 float64

	for {
		fmt.Println("Calculator Simulator")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Modulo")
		fmt.Println("6. Logarithm")
		fmt.Println("7. Square Root")
		fmt.Println("8. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter number 1: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter number 2: ")
			fmt.Scanln(&num2)
			result := add(num1, num2)
			fmt.Println("Result: ", result)
		case 2:
			fmt.Print("Enter number 1: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter number 2: ")
			fmt.Scanln(&num2)
			result := subtract(num1, num2)
			fmt.Println("Result: ", result)
		case 3:
			fmt.Print("Enter number 1: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter number 2: ")
			fmt.Scanln(&num2)
			result := multiply(num1, num2)
			fmt.Println("Result: ", result)
		case 4:
			fmt.Print("Enter number 1: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter number 2: ")
			fmt.Scanln(&num2)
			result, err := divide(num1, num2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Result: ", result)
			}
		case 5:
			fmt.Print("Enter number 1: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter number 2: ")
			fmt.Scanln(&num2)
			result, err := modulo(num1, num2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Result: ", result)
			}
		case 6:
			fmt.Print("Enter number: ")
			fmt.Scanln(&num1)
			result, err := logarithm(num1)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Result: ", result)
			}
		case 7:
			fmt.Print("Enter number: ")
			fmt.Scanln(&num1)
			result, err := squareRoot(num1)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Result: ", result)
			}
		case 8:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
