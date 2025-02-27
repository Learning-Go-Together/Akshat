package main

import "fmt"

/*
functions can also be passed as parameters as in GO function are datatype as int , float, bool
*/
func main() {
	fmt.Println(calculator(1, 2, 3, add))
	fmt.Println(calculator(1, 2, 3, subtract))
	fmt.Println(calculator(1, 2, 3, product))
	fmt.Println(calculator(1, 2, 3, divide))
}

func calculator(a, b, c int, operation func(a, b int) int) int {
	answer := operation(a, b)
	answer = operation(answer, c)
	return answer

}

func divide(a, b int) int {
	return a / b
}

func product(a, b int) int {
	return a * b
}

func subtract(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func add(a, b int) int {
	return a + b
}
