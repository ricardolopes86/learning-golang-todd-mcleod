package main

import (
	"fmt"
	"math"
)

// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

type Shape interface {
	area() float32
}

type Square struct {
	L float32
	W float32
}

type Circle struct {
	R float32
}

func (s Square) area() float32 {
	return s.L * s.W
}

func (s Circle) area() float32 {
	return s.R * math.Pi * 2
}

func info(t Shape) {
	fmt.Println(t.area())
}

func main() {

	c := Circle{
		R: 4,
	}

	s := Square{
		L: 5,
		W: 5,
	}
	info(c)
	info(s)

}
