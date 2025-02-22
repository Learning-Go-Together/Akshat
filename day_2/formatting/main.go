package main

import "fmt"

func main() {

	/*
		PRINTF -> used to print formatted Stringi in console
		PRINTLN -> used to print string with implicit "\n", we can only concate but not format
		PRINT -> can concat but not format, without "\n"
		SPRINT -> return string but not print it in console, but can concate it
		SPRINTF -> return formatted string, not print it
		SPRINTLN -> return formatted string with endline character "\n"


	*/
	age := 21
	fmt.Printf("My age is %d years.\n", age)

	salary := 1200.000000
	fmt.Printf("My salary is %f Rs.\n", salary)
	fmt.Printf("My salary is %.2f(round to 2 decimal places) Rs.\n", salary)

	s := fmt.Sprintf("My name is %v", "Akshat")
	fmt.Println(s)

}
