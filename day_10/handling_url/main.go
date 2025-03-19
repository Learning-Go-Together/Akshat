package main

import (
	"fmt"
	"net/url"
)

func main() {
	myurl := "https://randomuser.me/api/?results=5&callback=randomuserdata"
	parseUrl(myurl)
}

func parseUrl(myurl string) {
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println("---------------------")
	qparams := result.Query() //key value pairs of query params
	fmt.Println(qparams)
}
