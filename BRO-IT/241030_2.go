package main

// 30.10.2024

import "fmt"

func main() {
	//---------------------------------

	// 1) int, uint (8, 16, 32, 64)
	fmt.Println("(1)")
	var age uint8 = 23 // new int - age
	fmt.Println(age)   // output = age

	fmt.Println(" ")

	//---------------------------------

	// 2) float (32, 64)
	fmt.Println("(2)")
	var number float64 = 277.101 // new float - number
	fmt.Println(number)          // output = number

	fmt.Println(" ")

	//---------------------------------

	// 3) := - for auto type selecting
	fmt.Println("(3)")
	age1 := 16.3104   // same as var, but in this case Go will identify which type of data is age1 itself
	fmt.Println(age1) // output = age1

	fmt.Println(" ")

	//---------------------------------

	// 4) String + Length of string
	fmt.Println("(4)")
	name := "Aibek"        // new string - name
	fmt.Println(name)      // output = "Aibek"
	fmt.Println(len(name)) // output = 5 (length of our string

	fmt.Println(" ")

	//---------------------------------

	// 5) Input + Output some string with input
	fmt.Println("(5)")
	var name1 string // new empty string - name1
	fmt.Println("Your name:")
	fmt.Scan(&name1)                     // input data with the string type for name1
	fmt.Println("Hello, " + name1 + "!") // output = "Hello, (your string)!"

	fmt.Println(" ")

	//---------------------------------

	// 6) Transform int to string ||| (BONUS) Had some fun with if-else
	fmt.Println("(6) + (BONUS)")
	var age2 uint // new empty int - age2
	fmt.Println("Your age:")
	fmt.Scan(&age2) // input data with the integer type for age2

	if age2 < 18 { // "if" function, if the condition is true, then action inside will work
		fmt.Println(fmt.Sprint(age2) + "? You're too young") // in case with age2 less than 18, output = "(age2)? You're to young"
	} else { // "else" function, if "if" function is false, then action inside will work without any other checks
		fmt.Println("Nice!") // therefor, if age2 is greater or equal to 18, output = "Nice!"
	}

	//---------------------------------
}
