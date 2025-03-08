package main

import "fmt"

func main() {

	//variables declared here are at main package scope
	{
		//this is called block scope.
		//variables declared in this scope can only be accessible in this scope
		name := "Akshat"
		fmt.Println(name)

	}
	// fmt.Println(name)  <- this will give error bcoz name is used outside of it scope.
}
