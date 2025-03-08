package main

import "fmt"

func main() {
	fmt.Println(demo(6))
}

//defer is a keyword that called just before the termination of any function
func demo(a int) int {
	defer fmt.Println("Defer executed")
	if a == 3 {
		return a
	}

	return 0
}
