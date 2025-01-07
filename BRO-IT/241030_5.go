package main

// 30.10.2024

import "fmt"

func main() {
	//---------------------------------

	// 1) for{
	fmt.Println("(1)")
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, " + "Aibek10" + fmt.Sprint(i+1))
		fmt.Printf("What is that? %d\n", i+1)
	}

	fmt.Println(" ")

	//---------------------------------

	// 2) for{ (only even numbers)
	fmt.Println("(2)")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println(" ")

	//---------------------------------

	// 3) for{ + if{break}
	fmt.Println("(3)")
	for i := 1; i <= 100; i++ {
		fmt.Println(i)

		if i == 50 {
			break
		}
	}

	fmt.Println("That's it :)")
	fmt.Println(" ")

	//---------------------------------

	// 4) for == while
	fmt.Println("(4)")
	var a int
	fmt.Scan(&a)

	for a < 10 {
		fmt.Println(a)
		a++
	}

	fmt.Println(" ")

	//---------------------------------

	// 5) Fake range + array
	fmt.Println("(5)")
	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	fmt.Println(" ")

	//---------------------------------

	// 6) range + array
	fmt.Println("(6)")
	nums1 := []int{1, 2, 3, 4, 5}

	for index, element := range nums1 {
		fmt.Printf("Index: %d, Element: %d\n", index, element)
	}
	fmt.Println("\n")
	for index, _ := range nums1 {
		fmt.Printf("Index: %d\n", index)
	}
	fmt.Println("\n")
	for _, element := range nums1 {
		fmt.Printf("Element: %d\n", element)
	}

	fmt.Println(" ")

	//---------------------------------

	// 7) Two-dimensional array
	fmt.Println("(7)")
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}

		fmt.Println()
	}

	//---------------------------------
}
