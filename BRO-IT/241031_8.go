package main

// 31.20.2024

import (
	"fmt"
	"math"
)

func main() {
	//---------------------------------

	// 1) Mathematical operators "+", "-", "*", "/" ||| (BONUS) Calculator
	fmt.Println("(1) + (BONUS)")
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

	fmt.Println(" ")

	//---------------------------------

	// 2) Mathematical operator "%"
	fmt.Println("(2)")
	var a int
	var a1 int
	fmt.Println("Action is: %")
	fmt.Println("Enter a and b: ")
	fmt.Scan(&a, &a1)

	fmt.Println("Answer is: " + fmt.Sprint(a%a1))
	// "%"  returns only the remainder of the division

	fmt.Println(" ")

	//---------------------------------

	// 3) math.Sqrt()
	fmt.Println("(3)")
	var b float64
	fmt.Println("Action is: sqrt()")
	fmt.Println("Enter a: ")
	fmt.Scan(&b)

	answer := math.Sqrt(b)
	// Math.Sqrt() returns the square root of the number

	fmt.Println("Answer is: " + fmt.Sprint(answer))

	fmt.Println(" ")

	//---------------------------------

	// 4) math.Ceil()
	fmt.Println("(4)")
	var c float64
	fmt.Println("Action is: ceil()")
	fmt.Println("Enter a: ")
	fmt.Scan(&c)

	answer1 := math.Ceil(c)
	// Math.Ceil() returns the least int value greater than or equal to the number

	fmt.Println("Answer is: " + fmt.Sprint(answer1))

	fmt.Println(" ")

	//---------------------------------

	// 5) math.Floor()
	fmt.Println("(5)")
	var d float64
	fmt.Println("Action is: floor()")
	fmt.Println("Enter a: ")
	fmt.Scan(&d)

	answer2 := math.Floor(d)
	// Math.Floor() returns the greatest int value less than or equal to the number

	fmt.Println("Answer is: " + fmt.Sprint(answer2))

	fmt.Println(" ")

	//---------------------------------

	// 6) math.Round()
	fmt.Println("(6)")
	var e float64
	fmt.Println("Action is: round()")
	fmt.Println("Enter a: ")
	fmt.Scan(&e)

	answer3 := math.Round(e)
	// Math.Round() returns the nearest int

	fmt.Println("Answer is: " + fmt.Sprint(answer3))

	fmt.Println(" ")

	//---------------------------------

	// 7) Calculator
	fmt.Println("(7)")
	var f float64
	var f1 float64
	var action1 string
	fmt.Println("Select action: + - * /")
	fmt.Scan(&action1)
	fmt.Println("Enter a and b: ")
	fmt.Scan(&f, &f1)

	switch {
	case action1 == "+":
		fmt.Println("Answer is: " + fmt.Sprint(f+f1))
	case action1 == "-":
		fmt.Println("Answer is: " + fmt.Sprint(f-f1))
	case action1 == "*":
		fmt.Println("Answer is: " + fmt.Sprint(f*f1))
	case action1 == "/":
		fmt.Println("Answer is: " + fmt.Sprint(f/f1))
	}

	//---------------------------------
}
