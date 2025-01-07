package main

// 31.10.2024

import "fmt"

func main() {
	//---------------------------------

	// 1) Logical operators "and(&&)"
	fmt.Println("(1)")
	var a int
	var a1 int
	fmt.Scan(&a, &a1)

	if a > 0 && a1 > 0 {
		fmt.Println("+ & +")
	} else if a > 0 && a1 < 0 {
		fmt.Println("+ & -")
	} else if a < 0 && a1 > 0 {
		fmt.Println("- & +")
	} else if a < 0 && a1 < 0 {
		fmt.Println("- & -")
	} else {
		fmt.Println("NULL")
	}

	fmt.Println(" ")

	//---------------------------------

	// 2) Logical operator "or(||)"
	fmt.Println("(2)")
	var b int
	var b1 int
	fmt.Scan(&b, &b1)

	if b > 100 || b1 < -100 {
		fmt.Println("Out of range")
	} else {
		fmt.Println("You good!")
	}

	fmt.Println(" ")

	//---------------------------------

	// 3) Logical operator "not equal(!=)"
	fmt.Println("(3)")
	var c int
	fmt.Scan(&c)

	if c != 101 {
		fmt.Println("Access denied!")
	} else {
		fmt.Println("Welcome!")
	}

	fmt.Println(" ")

	//---------------------------------

	// 4) All together
	fmt.Println("(4)")
	d := 10
	d1 := 11

	if d != 0 && d1 < 100 || d > -100 && d1 != 0 {
		fmt.Println("TRUE!")
	} else {
		fmt.Println("FALSE!")
	}

	//---------------------------------
}
