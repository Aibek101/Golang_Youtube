package main

// 02.11.2024

import "fmt"

func main() {
	//---------------------------------

	// 1) Formatting strings (%s - string, %d - int, %f - float64, %t - bool)
	fmt.Println("(1)")
	age := 101
	name := "Aibek"
	money := 7777777.77
	check := true

	// fmt.Println("Age is: " + fmt.Sprint(age)) - before
	fmt.Printf("My name is: %s\nI'm: %d y.o.\nI have: %f$\nNo cap check: %t", name, age, money, check)

	//---------------------------------
}
