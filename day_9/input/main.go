package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')
	fmt.Println("Your age is :", age)
	fmt.Printf("Type is :%T\n", age)

}
