package main

import "fmt"

func main() {

	/*
		VARIABLES
		-> int8, int16, int32, int64
		-> uint8, uint16, uint32, uint64
		-> float32, float64
		-> complex32, complex64
		-> string
		-> bool

		%T -> placeholder used to get the type of variable

		variable declaration:

		1. using var -> Syntax: var <variable_name> <data_type>
		2. using walrus operator(:=)  -> can be only used inside any method , when the type is not confirmed
			and we want to assign that value Syntax: <variable_name> := <value>
		3. using only name ->  Syntax:  <variable_name> = <value>


		CONSTANTS
		We can create constant variables by adding "const" keyword :  const <var_name> <data_type>

		SAME LINE DECLARATION
		<var1_name>,<var2_name> :=  <value1>, <value2>


	*/
	x, y := 1, 2

	const pi float32 = 3.14
	const newConst = pi
	var normalInt int = 123423
	var unsignedInt uint = 234123
	var floatVal32 float32 = 234123.458945342
	var floatVal64 float64 = 234123.458945342

	fmt.Println(newConst)
	fmt.Println(x, y)
	fmt.Printf("Type :%T\n", x)
	fmt.Printf("Type :%T\n", y)
	fmt.Println(pi)
	fmt.Printf("Type :%T\n", pi)
	fmt.Println(normalInt)
	fmt.Printf("Type :%T\n", normalInt)
	fmt.Println(unsignedInt)
	fmt.Printf("Type :%T\n", unsignedInt)
	fmt.Println(floatVal32)
	fmt.Printf("Type :%T\n", floatVal32)
	fmt.Println(floatVal64)
	fmt.Printf("Type :%T\n", floatVal64)

}
