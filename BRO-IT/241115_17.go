package main

// 15.11.2024

import "fmt"

// (3)
func test(a func(int) int) {
	var a3 int
	fmt.Scan(&a3)

	fmt.Println(a(a3))
}

// (4)
func test1(x1 string) func() {
	return func() {
		fmt.Println(x1)
	}
}

func main() {
	//---------------------------------

	// 1) func() inside of main()
	fmt.Println("(1) - Input int, int. Output int+int")
	var a1, b1 int
	fmt.Scan(&a1, &b1)

	a := func(x int, y int) int {
		return x + y
	}

	sum := a(a1, b1)

	fmt.Println(sum)
	fmt.Println(" ")

	//---------------------------------

	// 2) func() inside of main()
	fmt.Println("(2) - Input int. Output sqr(int)")
	var a2 int
	fmt.Scan(&a2)

	square := func(a2 int) int {
		return a2 * a2
	}

	sqr := square(a2)

	fmt.Println(sqr)
	fmt.Println(" ")

	//---------------------------------

	// 3) func() that using another func()
	fmt.Println("(3) - Input int. Output sqr(int)")
	test(square)
	fmt.Println(" ")

	//---------------------------------

	// 4) return func() from func()
	fmt.Println("(4) - Input string. Output string x2")
	var a4 string
	fmt.Scan(&a4)

	// 1st option
	str := test1(a4)
	str()

	// 2nd option
	test1(a4)()

	//---------------------------------
}
