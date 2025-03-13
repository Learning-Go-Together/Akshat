package main

import (
	"fmt"
	"math"
)

func main() {

	var s1 shape = circle{2}
	var s2 shape = square{2}

	checkType(s1)
	checkType(s2)

}

type shape interface {
	area() float32
	perimeter() float32
}

type square struct {
	side float32
}

func (s square) area() float32 {
	return s.side * s.side
}

func (s square) perimeter() float32 {
	return s.side * 4
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

func checkType(t interface{}) {
	switch v := t.(type) {
	case square:
		fmt.Printf("%T <- square\n", v)
		fmt.Println(v.area())
		fmt.Println(v.perimeter())

	case circle:
		fmt.Printf("%T <- circle\n", v)
		fmt.Println(v.area())
		fmt.Println(v.perimeter())

	default:
		fmt.Println("Type not found")
	}
}
