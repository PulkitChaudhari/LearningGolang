package main

import "fmt"

func main() {
	var intSlice = []int{1,2,3}
	fmt.Println(sumSlice(intSlice))

	var floatSlice = []float32{1,2,3,4}
	fmt.Println(sumSlice(floatSlice))
}

func sumSlice[T int | float32](mySlice []T) T {
	var sum T
	for _, v := range(mySlice) {
		sum += v
	}
	return sum
}