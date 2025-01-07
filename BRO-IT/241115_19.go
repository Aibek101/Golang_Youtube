package main

// 15.11.2024 - Structures

import "fmt"

// (1)
type User struct {
	name     string
	age      int64
	password string
}

func change(u *User) {
	u.name = "Lebron"
	u.age = 39
	u.password = "Lakers"
}

func main() {
	//---------------------------------

	// 1) Structures
	fmt.Println("(1)")
	//var user User = User{name: "Aibek", age: 20, password: "123"}
	user := User{"Aibek", 20, "123"}
	user.name = "Aibek101"
	change(&user)

	fmt.Println(user)

	//---------------------------------
}
