package main

// 02.11.2024

import (
	"fmt"
	"sort"
)

func main() {
	//---------------------------------

	// 1) Slice + append() + sort.Ints()
	fmt.Println("(1)")
	slice := []int{3, 1, 2, 5, 7, 4}
	slice = append(slice, 0) // adding new element in our slice
	// "slice[0] = 11" will replace first element(3) with "11"
	sort.Ints(slice)

	fmt.Println(slice)

	fmt.Println(" ")

	//---------------------------------

	// 2) Slice + sort.Strings()
	fmt.Println("(2)")
	slice1 := []string{"LeBron James", "Leonel Messi", "Jackie Chan", "Aibek101"}
	sort.Strings(slice1)

	fmt.Println(slice1)

	//---------------------------------
}
