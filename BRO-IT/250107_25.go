package main

// 07.01.2025 - goto

import (
	"fmt"
)

func main() {
	//---------------------------------

	// 1)
	fmt.Println("(1) goto")
	i := 0

LOOP:
	if i > 50 {
		goto END
	}
	fmt.Println(i)
	i++
	goto LOOP

END:
	fmt.Println("END")

	fmt.Println(" ")

	//---------------------------------
}
