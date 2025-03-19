package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://www.wikipedia.org/"
	getRequest(url)
}

func getRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type :%T", response)
	defer response.Body.Close()

	databytes, er := io.ReadAll(response.Body)
	if er != nil {
		panic(er)
	}

	content := string(databytes)
	fmt.Println(content)

}
