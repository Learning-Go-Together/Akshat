package main

import "fmt"

func main() {
	fmt.Println(naked_return())

}

//this is naked return if value for s is not specified then it returns default value
func naked_return() (s string) {
	s = "ask"
	return
}
