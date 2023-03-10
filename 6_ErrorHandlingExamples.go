package main

import (
	"errors"
	"fmt"
	"math"
)

func getHydroQuebecBill(lastYearPrice, thisYearPrice float64) (float64, float64, error) {
	if lastYearPrice == 0 {
		err := errors.New("HydroQuebec bill can never be 0!")
		return 0, 0, err
	}

	difference := lastYearPrice - thisYearPrice
	billRatio := (difference / lastYearPrice) * 100
	return difference, billRatio, nil
}

func main() {
	lastYearPrice := 0.00
	thisYearPrice := 170.0

	difference, billRatio, err := getHydroQuebecBill(lastYearPrice, thisYearPrice)

	if err != nil {
		fmt.Println("Sorry! An exeption happened!")
	} else {
		if difference < 0 {
			fmt.Printf("The bill is down by $%.2f relative to last year $%.2f", math.Abs(difference), math.Abs(billRatio))
		} else {
			fmt.Printf("The bill is up by $%2.f relative to last year $%2.f", math.Abs(difference), math.Abs(billRatio))
		}
	}

}
