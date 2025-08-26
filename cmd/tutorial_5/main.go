package main

import "fmt"

func main() {
	var ptr1 *int32 = new(int32)
	var ptr2 int32 = 32
	fmt.Println(ptr1)
	fmt.Println(*ptr1)
	fmt.Println(ptr2)
	fmt.Println(&ptr2)

	var mySlice = []int32{1,2,3}
	var myCopySlice = mySlice;
	mySlice[2] = 100
	fmt.Println(mySlice[2], myCopySlice[2])

	var arr = [5]int32{1,2,3,4,5}
	var temp = squareArray(arr)
	fmt.Println(temp)
	fmt.Println(arr)

}

func squareArray(inputArray [5]int32) [5]int32 {
	for i := range(inputArray) {
		inputArray[i] = inputArray[i] * inputArray[i]
	}
	return inputArray
}