package main

import (
	"fmt"
)

func main() {

	// Using strings to iterate
	var myString string = "Résuméwdss"
	for i := range(myString) {
		fmt.Println(i, " ", myString[i])
	}

	// Casting strings to runes to iterate
	var myRune = []rune("Résuméwdss")
	for i := range(myRune) {
		fmt.Println(i, " ", myRune[i])
	}
}