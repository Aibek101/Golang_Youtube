package main

// 29.11.2024 - Interfaces

import "fmt"

type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	sum() int
	sub() int
	mul() int
	div() float64
}

func (n Numbers) sum() int {
	return n.num1 + n.num2
}
func (n Numbers) sub() int {
	return n.num1 - n.num2
}
func (n Numbers) mul() int {
	return n.num1 * n.num2
}
func (n Numbers) div() float64 {
	return float64(n.num1) / float64(n.num2)
}

func main() {
	//---------------------------------

	// 1)
	fmt.Println("(1) Calculator")
	var a, b int
	fmt.Print("Enter the numbers: ")
	fmt.Scan(&a, &b)

	var i NumbersInterface
	numbers := Numbers{a, b}
	i = numbers

	fmt.Printf("Total = %d\n", i.sum())
	fmt.Printf("Difference = %d\n", i.sub())
	fmt.Printf("Product = %d\n", i.mul())
	fmt.Printf("Quotient = %f\n", i.div())

	//---------------------------------
}
