package main

// 07.01.2025 - Goroutines

import (
	"fmt"
)

func main() {
	//---------------------------------

	// 1)
	fmt.Println("(1) Goroutines")
	ch := make(chan string)
	ch2 := make(chan int)

	go say("Hello Go!", ch, ch2)

	fmt.Println(<-ch, <-ch2)

	fmt.Println(" ")

	//---------------------------------

	// 2)
	fmt.Println("(2) Goroutines")
	ch3 := make(chan int)

	go say1("Hello Go!", ch3)

	for a := range ch3 {
		fmt.Println(a)
	}

	//---------------------------------
}

// 1)
func say(greet string, ch chan string, ch2 chan int) {
	fmt.Println(greet)
	ch <- "Channel"
	ch2 <- 101
}

// 2)
func say1(greet string, ch3 chan int) {
	for i := 0; i <= 5; i++ {
		ch3 <- i
	}

	close(ch3)
}
