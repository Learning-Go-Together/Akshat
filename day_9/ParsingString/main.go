package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to voting System!")
	fmt.Print("Enter your age:")
	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')

	//since the input is of string type we need to parse it into int type
	// num, err := strconv.ParseInt(inputStr, 10, 64) //0 strconv.ParseInt: parsing "23\r\n": invalid syntax
	age, err := strconv.ParseInt(strings.TrimSpace(inputStr), 10, 64) //Trim will remove all the spaces

	if age > 18 && age < 100 {
		fmt.Println("You are eligible to vote!")
	} else if age < 18 && age > 0 {
		fmt.Println("Not eligible to vote!")
	} else if age <= 0 {
		fmt.Println("age can't be negative!")
	} else {
		fmt.Println(err)
	}

}
