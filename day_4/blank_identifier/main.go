package main

import "fmt"

func main() {
	//by using blank identifier "_" we can ignore the values that we dont need
	sum, _ := add(2, 3)
	fmt.Println(sum)
}

func add(a int, b int) (int, int) {
	return a + b, -1

}
