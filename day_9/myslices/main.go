package main

import (
	"fmt"
	"sort"
)

func main() {
	var sl = []string{"hello"}
	sl = append(sl, "this", "is", "slice")
	fmt.Println(sl)
	fmt.Println(len(sl))

	fmt.Println(sl[1:])
	fmt.Println(sl[1:3])
	sl = sl[1:3]
	fmt.Println(sl)

	//another way of declaring slices

	newSlice := make([]int, 5)
	newSlice[0] = 1
	newSlice[2] = 1
	newSlice = append(newSlice, 45)
	fmt.Println(newSlice)

	fmt.Println(sort.IntsAreSorted(newSlice))
	sort.Ints(newSlice)
	fmt.Println(sort.IntsAreSorted(newSlice))
}
