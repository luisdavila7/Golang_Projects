package main

import "fmt"

func Reversal(s *string) {
	runes := []rune(*s)

	// REseach about this
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*s = string(runes)
}

func main() {
	srt := "Dale Duro"

	fmt.Println("Original phrase", srt)
	Reversal(&srt)
	fmt.Println("Reverse phrase", srt)
}

