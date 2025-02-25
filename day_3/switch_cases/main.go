package main

import "fmt"

/*
In switch cases we dont explicity need to add a break statement
but to fall through we have to explicitly do it with "fallthrough" keyword
*/

func main() {
	var signal string = "red"
	switch signal {
	case "red":
		fmt.Println("Stop!")
	case "yellow":
		fmt.Println("Be ready!")
	case "green":
		fmt.Println("Go!")
	default:
		fmt.Println("Bhad me jao!")
	}

	var n int = 1
	switch n {
	case 0:
		fmt.Println("Smaller than 0")
	case 1:
		fmt.Println("larger than 0")
		fallthrough
	case 2:
		fmt.Println("larger than 0")
		fallthrough
	default:
		fmt.Println("n is a number")

	}

}
