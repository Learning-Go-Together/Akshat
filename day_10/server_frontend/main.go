package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	posturl := "http://localhost:8080/public/create-user"
	// geturl := "http://localhost:8080/journal"
	// getResponse(geturl)

	// postJsonResponse(posturl)
	postFormResponse(posturl)
}

func getResponse(myurl string) {
	response, err := http.Get(myurl)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	databytes, er := io.ReadAll(response.Body)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(string(databytes))

}

func postJsonResponse(myurl string) {
	body := strings.NewReader(
		`
		{
			"username":"demoUser",
			"password":"1234"
		}
		`)
	response, err := http.Post(myurl, "application/json", body)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("body: ", response.Body)
}

func postFormResponse(myurl string) {
	data := url.Values{}
	data.Add("username", "newUser")
	data.Add("password", "newUser")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

}
