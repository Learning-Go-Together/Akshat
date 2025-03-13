package main

import "fmt"

func main() {

	var v1 vehicle = maruti{4, 5}
	fmt.Println(v1.(vehicle))

	var v2 vehicle = maruti{4, 7}
	b, ok := v2.(Bolero)
	fmt.Println(b, ok)

}

type vehicle interface {
	start()
}

type car interface {
	blowHorn()
}

type maruti struct {
	wheels int
	seats  int
}

type Bolero struct {
	maruti
}

func (m maruti) start() {
	fmt.Println("Maruti Started!")
}

func (m maruti) blowHorn() {
	fmt.Println("Horn Blown!")
}
