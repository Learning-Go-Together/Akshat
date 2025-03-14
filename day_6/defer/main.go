package main

import "fmt"

func main() {
	fmt.Println(demo(6))
}

//defer is a keyword that called just before the termination of any function

//if there are multiple defer statements then they are executed in reverse order
func demo(a int) int {
	defer fmt.Println("Defer executed")
	if a == 3 {
		return a
	}

	return 0
}
