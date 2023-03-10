package main

import (
	"fmt"
	"math"
)

func compareHydroQuebecBill(lastYearPrice, thisYearPrice float64) (float64, float64) {
	difference := thisYearPrice - lastYearPrice
	billRatio := (difference / lastYearPrice) * 100
	return difference, billRatio
}

func main() {

	lastYearPrice := 0.0
	thisYearPrice := 160.0

	difference, billRatio := compareHydroQuebecBill(lastYearPrice, thisYearPrice)

	if difference < 0 {
		fmt.Printf("The bill is down by $%.2f relative to last year $%.2f", math.Abs(difference), math.Abs(billRatio))
	} else {
		fmt.Printf("The bill is up by $%.2f relative to last year $%.2f", difference, billRatio)
	}
}
