package main

// 02.11.2024

import "fmt"

func main() {
	//---------------------------------

	// 1) Multidimensional array
	fmt.Println("(1)")
	arr := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(arr)
	fmt.Println(arr[1][2]) // Should be "7"; 1 - second row, 2 - third column

	//---------------------------------
}
