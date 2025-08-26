package main

import (
	"fmt"
)

var pi float32 = 3.14

type shape interface {
	getArea() float32
}

type circle struct {
	radius int
}

type square struct {
	side int
}

func (e square) getArea() float32 {
	return float32(e.side * e.side)
}

func (e circle) getArea() float32 {
	return pi * float32(e.radius * e.radius)
}

func isHuge(e shape) bool {
	return e.getArea() > 100.0
}

func main() {
	var myCircle circle = circle{1}
	fmt.Println(isHuge(myCircle))

	var mySquare square = square{11}
	fmt.Println(isHuge(mySquare))
}