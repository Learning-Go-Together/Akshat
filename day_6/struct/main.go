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
	// alto.demo4()
	demo4(alto)
	alto.demo5()
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

//passing struct as arguments
func demo4(ca car) {
	fmt.Println(ca.brand, ca.model, ca.regNo)
}

//struct methods
func (c car) demo5() {
	fmt.Println("Struct methods called")
}

//structs fields are stored in same manner as
/*	struct stores all the field in contigous manner so there order matters as it divides
	the memory according to the largest datatype


	type stats struct {
		Reach    uint16  <- 2 bytes
		NumPosts uint8 	 <- 1 bytes
		NumLikes uint8	 <- 1 bytes
	}

	memory :
	   2	  2
	[  2  ][ 1 ,1 ]


	type stats struct {
		NumPosts uint8 	 <- 1 bytes
		Reach    uint16  <- 2 bytes
		NumLikes uint8	 <- 1 bytes
	}

	memory:
		2		2		2
	[	1	][	2	][	1	]
*/

//empty struct

func demoEmptyStruct() {

	//anonymous empty struct
	// emptyStruct := struct{}{}
	//empty struct declaration
	type emptyStruct struct{}
}
