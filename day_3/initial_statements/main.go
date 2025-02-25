package main

import "fmt"

func main() {
	/*
		initial statements is a way to declare any variable or write any statement that is meant
		to be in the if scope and not in parent scope.

		Syntax :

		if <initial_statement> ; <condition>{
			#####
		}
	*/

	if len := 4; len > 0 {
		fmt.Println("Yes its greater!")
	} else {
		fmt.Println("No its not!")
	}
}
