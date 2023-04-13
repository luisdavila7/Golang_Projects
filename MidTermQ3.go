package main

import "fmt"

func validatePrimeNumber(prime int) bool {
	if prime <= 1 {
		return false
	}
	for i := 2; i*i <= prime; i++ {
		if prime%i == 0 {
			return false
		}
	}
	return true
}

func arrayPrimeNum(arrPrime []int) []bool {
	result := make([]bool, len(arrPrime))
	for i, num := range arrPrime {
		result[i] = validatePrimeNumber(num)
	}
	return result
}

func main() {

	arrPrime := []int{2, 4, 29, 8, 32}
	isPrime := arrayPrimeNum(arrPrime)
	fmt.Println(isPrime)
}
