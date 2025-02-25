package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2

	fmt.Println(add(a, b))
	fmt.Println(subtract(a, b))
}

func add(a int, b int) int {
	return a + b

}

// we can write parameter with same type together like this
func subtract(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
