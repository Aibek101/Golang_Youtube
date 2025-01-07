package main

// 02.11.2024

import "fmt"

func main() {
	//---------------------------------

	// 1) Map
	fmt.Println("(1)")
	var a string
	fmt.Scan(&a)

	money := map[string]int{
		"USD": 5000000,
		"KZT": 10000000,
		"EUR": 40000,
	}
	// "money["USD"] = 30000000" - change
	// "delete(money, "USD")" - delete
	// el, status := money["USD"]; el = 5000000, status = true(we have "USD" in map)

	fmt.Printf("Your balance is: %d", money[a])

	//---------------------------------
}
