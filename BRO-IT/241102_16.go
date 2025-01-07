package main

// 02.11.2024, 15.11.2024

import "fmt"

// ---------------------------------

// 1) Function
func first() {
	fmt.Println("Hello from function")
}

//---------------------------------

// 2) Function - int

func sum(x1 int, y1 int) int {
	result := x1 + y1

	return result
}

//---------------------------------

// 3) Function
func math_func(x2 int, y2 int) (sum1 int, sub1 int, mul1 int, div1 int) {
	sum1 = x2 + y2
	sub1 = x2 - y2
	mul1 = x2 * y2
	div1 = x2 / y2

	return
}

//---------------------------------

func main() {
	fmt.Println("(1)")
	first()
	fmt.Println(" ")

	fmt.Println("(2)")
	var x1, y1 int
	fmt.Scan(&x1, &y1)
	value := sum(x1, y1)
	println(value)
	fmt.Println(" ")

	fmt.Println("(3)")
	var x2, y2 int
	fmt.Scan(&x2, &y2)
	result1, result2, result3, result4 := math_func(x2, y2)
	fmt.Println(result1, result2, result3, result4)
}
