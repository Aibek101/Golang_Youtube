package main

// 07.01.2025 - Files with ioutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//---------------------------------

	// 1)
	fmt.Println("(1) Files with ioutil")
	data := []byte("Text to file")
	e := ioutil.WriteFile("text22.txt", data, 0600)
	if e != nil {
		fmt.Println("Can not create a file\n", e)
	}

	file_data, err := ioutil.ReadFile("text22.txt")
	if err != nil {
		fmt.Println("I can't read text.txt\n", err)
	}
	fmt.Println(string(file_data))

	os.Remove("text22.txt")

	//---------------------------------
}
