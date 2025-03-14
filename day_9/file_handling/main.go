package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter filename:")
	filename, err := reader.ReadString('\n')
	fmt.Println("Enter content:")
	content, er := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	if er != nil {
		fmt.Println(er)
	}

	createFile(strings.TrimSpace(filename), strings.TrimSpace(content))
	readAFile(strings.TrimSpace(filename))

}
func readAFile(filename string) {
	databytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(databytes) //this will print databytes
	fmt.Println(string(databytes))

}

func createFile(filename, content string) {
	filename = "./" + filename
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	length, er := io.WriteString(file, content)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println("Length is:", length)
}
