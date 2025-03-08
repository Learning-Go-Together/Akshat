package main

import "fmt"

func main() {
	var alto car
	//or
	// alto := car{}
	alto.brand = "Maruti"
	alto.model = "alto"
	alto.regNo = 12

	// var splendor bike
	splendor := bike{}
	splendor.brand = "Hero"
	splendor.mileage = 70
	splendor.model = "splendor+"
	splendor.backWheel.material = "Alloy"
	splendor.frontWheel.material = "Alloy"
	fmt.Println("Car: ", alto)
	fmt.Println("Bike: ", splendor)
	demo()
	demo2()
	demo3()
}

type car struct {
	brand string
	model string
	regNo int
}

//Nested Structs

type bike struct {
	brand      string
	model      string
	mileage    int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

// Anonymous Struct
func demo() {

	demoStruct := struct {
		id   int
		name string
	}{
		id:   12,
		name: "Akshat",
	}

	fmt.Println(demoStruct)
}

//Nested anonymous struct

func demo2() {

	type bike struct {
		brand   string
		model   string
		mileage int
		wheel   struct {
			radius   int
			material string
		}
	}

	bikes := bike{
		brand:   "honda",
		model:   "cbShine",
		mileage: 50,
		wheel:   wheel{12, "Alloy"},
	}

	fmt.Println(bikes)
}

//embedded structs
func demo3() {
	type car struct {
		brand string
		model string
	}

	type truck struct {
		// "car" is embedded, so the definition of a
		// "truck" now also additionally contains all
		// of the fields of the car struct
		car
		bedSize int
	}

	demoStruct := truck{
		car:     car{"maruti", "alto"},
		bedSize: 2,
	}

	fmt.Println(demoStruct)

}
