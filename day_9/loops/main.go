package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//for each loop
	for idx, val := range nums {
		fmt.Println(idx, val)
	}

	for i := range 5 {
		fmt.Println(i)
	}

	//for loop
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
	//while loop
	value := 6

	for value <= 12 {
		fmt.Println(value)
		value++
	}
}
