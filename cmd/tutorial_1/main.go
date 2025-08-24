package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var quotient, remainder, err = divisionTool(100,7)
	if err != nil {
		fmt.Printf("Division not possible because, %v", err.Error())
	} else {
		fmt.Printf("When 100 is divided by 7, the quotient is %v and remainder is %v",quotient, remainder)	
	}
}

func divisionTool(numerator int, denominator int) (int, int, error) {
	var err error	
	if denominator == 0 {
		err = errors.New("Denominator is zero")
		return 0, 0, err
	}
	quotient := numerator / denominator
	remainder := numerator % denominator
	return quotient, remainder, err
}
