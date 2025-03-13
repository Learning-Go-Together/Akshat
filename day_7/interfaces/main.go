package main

import (
	"fmt"
	"math"
)

func main() {

	s := square{2}
	printShapeData(s)

	c := circle{2.5}
	printShapeData(c)

}

// since we provided the implementation of area() method thus we can use it and it calls the implementation of s shape
func printShapeData(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

type shape interface {
	area() float32
	perimeter() float32
}

type square struct {
	side float32
}

//Here struct square has implemented the interface method so its an implementation of interface

// if its implements any of the method then it has to implement all the methods
func (s square) area() float32 {
	return s.side * s.side
}

func (s square) perimeter() float32 {
	return s.side * 4
}

// since this struct doesn't implemented any of the interface method so it is not an implementation of interface
type rectangle struct {
	length  float32
	breadth float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c circle) getArea() float32 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}
