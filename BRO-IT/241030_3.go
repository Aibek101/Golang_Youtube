package main

// 30.10.2024

import (
	"fmt"
	"math"
)

func main() {
	//---------------------------------

	// 1) if{, else{, else if{
	fmt.Println("(1)")
	var num int
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("Your number is greater than 0")
	} else if num == 0 {
		fmt.Println("Your number is 0")
	} else {
		fmt.Println("Your number is less than 0")
	}

	fmt.Println(" ")

	//---------------------------------

	// 2) Program that will solve a quadratic equations
	fmt.Println("(2)")
	var a float64
	var b float64
	var c float64

	fmt.Println("\nQuadratic equation calculator")
	fmt.Println("Enter a:")
	fmt.Scan(&a)
	fmt.Println("Enter b:")
	fmt.Scan(&b)
	fmt.Println("Enter c:")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Answer is: " + fmt.Sprint(x1) + ", " + fmt.Sprint(x2))
	} else if D == 0 {
		var x0 float64
		x0 = (-b) / (2 * a)

		fmt.Println("Answer is: " + fmt.Sprint(x0))
	} else {
		fmt.Println("The equation has no solution")
	}

	fmt.Println(" ")

	//---------------------------------

	// (BONUS) Calculator
	fmt.Println("(BONUS)")
	var num1 float64
	var num2 float64
	var action string

	fmt.Println("\nCalculator")
	fmt.Println("Enter a: ")
	fmt.Scan(&num1)
	fmt.Println("Select action: + - * /")
	fmt.Scan(&action)
	fmt.Println("Enter b: ")
	fmt.Scan(&num2)

	if action == "+" {
		fmt.Println("Answer is: " + fmt.Sprint(num1+num2))
	} else if action == "-" {
		fmt.Println("Answer is: " + fmt.Sprint(num1-num2))
	} else if action == "*" {
		fmt.Println("Answer is: " + fmt.Sprint(num1*num2))
	} else if action == "/" {
		fmt.Println("Answer is: " + fmt.Sprint(num1/num2))
	}

	//---------------------------------
}

// line 15 :: new empty int - num
// line 16 :: input data with the int type for num
// line 18 :: "if" function, num greater than 0
// line 19 :: output = "Your number is greater than 0"
// line 20 :: "else if" or "elif(another programming language)" function, with this function you're able to add another condition to check, num is equal to 0
// line 21 :: output = "Your number is 0"
// line 22 :: "else" function, num less than 0
// line 23 :: output = "Your number is less than 0"
// line 32 :: new empty float - a
// line 33 :: new empty float - b
// line 34 :: new empty float - c
