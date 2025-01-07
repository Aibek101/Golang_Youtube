package main

// 15.11.2024

import "fmt"

// (3)
func change(str *string) {
	*str = "GODLIKE"
}

func main() {
	//---------------------------------

	// 1) Pointers (&)
	fmt.Println("(1) - Input int. Output 'Address in the RAM'")
	var a int
	fmt.Scan(&a)

	pointer1 := &a

	fmt.Println(pointer1)
	fmt.Println(" ")

	//---------------------------------

	// 2) Pointers (*)
	fmt.Println("(2) - Input int. Output int")
	var b int
	fmt.Scan(&b)

	pointer2 := &b

	fmt.Println(*pointer2)
	fmt.Println(" ")

	//---------------------------------

	// 3) Using pointers in func()
	fmt.Println("(3) - Input string. Output changed string")
	var s string
	fmt.Scan(&s)
	fmt.Println("Your input: " + s)
	var pointer3 *string = &s
	change(pointer3)
	fmt.Println("Your changed input: " + s)
	fmt.Println(" ")

	//---------------------------------

	// 4)
	fmt.Println("(4) - Input int. Output int + 'Address in the RAM'")
	var c int
	fmt.Scan(&c)

	c1 := &c

	fmt.Println(c, c1)

	//---------------------------------
}
