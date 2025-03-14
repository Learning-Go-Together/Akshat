package main

import "fmt"

func main() {
	fmt.Println("hello")
	defer fmt.Println("World")
	demodefer()
}

func demodefer() {
	for i := range 3 {
		defer fmt.Println(i)
	}
}
