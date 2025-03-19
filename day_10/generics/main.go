package main

import "fmt"

func main() {
	intslice := []int{1, 2, 3, 4, 5, 6}
	stringslice := []string{"hkndf", "kjsbfk", "sjdbf"}

	fmt.Println(divideSlice(intslice))
	fmt.Println(divideSlice(stringslice))
}

//generic function

func divideSlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}
