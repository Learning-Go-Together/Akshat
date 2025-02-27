package main

import "fmt"

func main() {
	fmt.Println(naked_return())
	fmt.Println(explicit_return())

}

//this is naked return if value for s is not specified then it returns default value
func naked_return() (s string) {
	s = "ask"
	return
}

//this is explicit return
func explicit_return() (s string) {
	return "explicit"
}

