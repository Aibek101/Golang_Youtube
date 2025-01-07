package main

// 30.10.2024

import (
	"fmt"
	"math"
)

func main() {
	//---------------------------------

	// 1) go build main.go + go build -ldflags "-s -w" main.go
	fmt.Println("(1)")
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

	fmt.Scan()

	//---------------------------------
}
