package main

// 31.10.2024 - 02.11.2024

import (
	"fmt"
	"math"
)

func main() {
	//---------------------------------

	// 1) Array - string
	fmt.Println("(1)")
	names := [3]string{"Aibek", "LeBron", "Messi"}
	fmt.Println(names)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println(" ")

	//---------------------------------

	// 2) Array - int
	fmt.Println("(2)")
	marks := [5]float64{11, 9, 12, 8, 10}
	// better to use float64 not int, to have no problems with numbers in the future, for example if avg will be 6.36 instead of 6
	// for "fmt.Println(marks[2])" output will be "12"
	var sum float64 = 0

	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}

	var avg float64 = sum / float64(len(marks))
	fmt.Println("Answer is: " + fmt.Sprint(math.Round(avg)))

	//---------------------------------
}
