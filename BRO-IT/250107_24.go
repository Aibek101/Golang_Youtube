package main

// 07.01.2025 - Files with os

import (
	"fmt"
	"os"
)

func main() {
	//---------------------------------

	// 1)
	fmt.Println("(1) Files with os")
	file, err := os.Create("text24.txt")
	if err != nil {
		fmt.Println("ERROR -", err)
		os.Exit(1)
	}

	defer file.Close()

	data := "Ion\n\nLeshrac\nOwl\nVendeta\nEvangelion\n\nYegger\nOil\nUranium"
	file.WriteString(data)

	file_data, err1 := os.ReadFile(file.Name())
	if err1 != nil {
		fmt.Println("ERROR -", err1)
		os.Exit(1)
	}

	fmt.Println(string(file_data))

	os.Remove("text24.txt")

	//---------------------------------
}
