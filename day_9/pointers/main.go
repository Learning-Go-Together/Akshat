package main

import "fmt"

func main() {
	myNumber := 12
	var ptr *int = &myNumber

	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr + 5
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
