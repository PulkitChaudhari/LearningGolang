package main

import (
	"fmt"
)

func main() {
	var myArray [3]int = [3]int{1,2,3}
	fmt.Println(myArray)
	fmt.Println(myArray[1:3])
	fmt.Println(myArray[1])

	myOtherArray := [...]int{1,2,3,4,5,6}
	fmt.Println(myOtherArray)


	// Slices

	var mySlice []int = []int{4,5,6}
	fmt.Println(mySlice)
	fmt.Printf("Length of Slice is %v and Capacity of Slice is %v \n", len(mySlice), cap(mySlice))
	mySlice = append(mySlice,7)
	fmt.Println(mySlice)
	fmt.Printf("Length of Slice is %v and Capacity of Slice is %v \n ", len(mySlice), cap(mySlice))

	var tempSlice []int = []int {1,2,3}

	tempSlice = append(tempSlice,mySlice...)
	fmt.Println(tempSlice)

	var myCustomSlice []int = make([]int, 5 , 10)
	fmt.Printf("Length of Slice is %v and Capacity of Slice is %v \n", len(myCustomSlice), cap(myCustomSlice))

	// Maps

	var myEmptyMap map[string]int = make(map[string]int)
	fmt.Println(myEmptyMap)

	var myMap map[string]int = map[string]int{"Jason":23,"Bourne":32}
	fmt.Println(myMap)
	fmt.Println(myMap["Jason"])
	var nameFromMap, ageFromMap = myMap["Jason"]
	fmt.Println(nameFromMap, ageFromMap)

	// Loops

	for name, age := range(myMap) {
		fmt.Printf("Name is %v and Age is %v \n", name, age)
	}

	for i := 0; i < len(myArray); i++ {
		fmt.Printf("Element of myArray at %v index is : %v \n", i, myArray[i])
	}
}