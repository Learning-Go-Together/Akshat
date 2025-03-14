package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "orange", "peach"}
	//suppose we want to remove the element at index 2 i.e orange
	fmt.Println(fruits)
	fruits = append(fruits[:2], fruits[3:]...)
	fmt.Println(fruits)

}
