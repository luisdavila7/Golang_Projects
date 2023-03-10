package main

import "fmt"

func areaRectangle(width float64, height float64) float64 {
	return (width * height)
}

func main() {
	width := 8.90
	height := 12.07
	result := areaRectangle(width, height)
	fmt.Printf("The area with width %.2f and height %2f are %2.f", width, height, result)
}
