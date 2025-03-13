package main

import "fmt"

type add interface {
	addTwoNum() int
	// addThreeNum() int

}

func calculateDetails(a add, s subtract) {
	fmt.Println(a.addTwoNum())
	fmt.Println(s.subTwoNum())
}

type subtract interface {
	subTwoNum() int
}

type calculator struct {
	num1 int
	num2 int
}

func (c calculator) addTwoNum() int {
	return c.num1 + c.num2
}
func (c calculator) subTwoNum() int {
	if c.num1 > c.num2 {
		return c.num1 - c.num2
	}
	return c.num2 - c.num1
}

func main() {
	c := calculator{1, 2}
	calculateDetails(c, c)

}
