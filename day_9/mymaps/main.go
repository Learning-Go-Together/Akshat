package main

import "fmt"

func main() {
	//creating a map

	languageMap := make(map[string]string)

	languageMap["js"] = "javascript"
	languageMap["py"] = "python"
	languageMap["rb"] = "ruby"

	fmt.Println(languageMap)

	delete(languageMap, "py")
	fmt.Println(languageMap)

	for key, val := range languageMap {
		fmt.Println(key, val)
	}
}
