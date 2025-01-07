package main

// 02.11.2024

import "fmt"

func main() {
	//---------------------------------

	// 1) Slice + range
	fmt.Println("(1)")
	slice := []int{1, 4, 6, 2, 3, 9, 3}

	for _, el := range slice {
		fmt.Printf("%d\n", el)
	}

	//---------------------------------
}
