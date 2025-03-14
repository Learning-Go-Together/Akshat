package main

import "fmt"

func main() {
	var fruits [5]string
	fruits[0] = "apple"
	fruits[2] = "banana"
	fruits[3] = "pear"
	fruits[4] = "papaya"

	fmt.Println(fruits)      //this will add extra space at index 1 as no value is provided
	fmt.Println(len(fruits)) // this will return the max size of array no matter of the no. of values

	var vegatables = [2]string{"tomato", "brinjal"}
	fmt.Println(vegatables)
}
