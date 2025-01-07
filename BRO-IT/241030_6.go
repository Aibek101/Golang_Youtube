package main

// 30.10.2024

import "fmt"

func main() {
	//---------------------------------

	// 1) Switch
	fmt.Println("(1)")
	var name string
	fmt.Scan(&name)

	switch name {
	case "Aibek":
		fmt.Println("Greetings my lord!")
	case "U":
		fmt.Println("Greetings my queen!")
	default:
		fmt.Println("Who the f*ck are you?!")
	}

	fmt.Println(" ")

	//---------------------------------

	// 2) Switch (fallthrough)
	fmt.Println("(2)")
	var num int
	fmt.Scan(&num)

	switch {
	case num > 0:
		fmt.Println("Number is greater than 0")
		fallthrough
	case num < 100:
		fmt.Println("Number is less than 100")
		fallthrough
	case num <= 0:
		fmt.Println("Number is equal or less than 0")
	}
	// fallthrough will run next case without conditions checking

	//---------------------------------
}
