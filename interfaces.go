package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface { //common : grps types together based on methods
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

//TO BE PART OF SHAPE INTERFACE U NEED TO HAVE both AREA AND CIRCUMF
//FUNC ASSOCIATED WITH YOU

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// take in type shape
func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	//u can create an slice of type shape now as well
	shapes := []shape{
		square{length: 15.2}, //struct types
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
