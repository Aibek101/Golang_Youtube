package main

// 15.11.2024, 29.11.2024 - Advanced Structures

import "fmt"

type User struct {
	name     string
	age      int64
	password string
	bonus    []int
}

// (1) ---
func (u *User) setName(name1 string) {
	u.name = name1
}

// (2)
func (u User) isElder() bool {
	a := u.age
	isTrue := false

	if a >= 18 {
		isTrue = true
	} else if a < 18 {
		isTrue = false
	}

	return isTrue
}

// (3)
func (u User) getHighestBonus() int {
	max := 0

	for _, sc := range u.bonus {
		if max < sc {
			max = sc
		}
	}

	return max
}

func main() {
	var name101 string
	var age101 int64
	var password101 string
	fmt.Print("Enter the NAME, AGE and PASSWORD: ")
	fmt.Scan(&name101, &age101, &password101)
	//var user User = User{name101, age101, password101}
	user := User{name101, age101, password101, []int{101, 23, 36, 28, 76, 22498, 69}}
	fmt.Println(user)
	fmt.Println(" ")
	//---------------------------------

	// 1) Structures
	fmt.Println("(1) Changing the name")
	var newName string
	fmt.Print("New name: ")
	fmt.Scan(&newName)
	user.setName(newName)

	fmt.Println(user)

	fmt.Println(" ")

	//---------------------------------

	// 2) Age check
	fmt.Println("(2) Checking the age(18+)")
	if user.isElder() {
		fmt.Println("You can use this site :)")
	} else {
		fmt.Println("You are not ready!")
	}
	fmt.Println(" ")

	//---------------------------------

	// 3) Highest Bonus Earned
	fmt.Println("(3) Checking the highest bonus earned")
	fmt.Printf("Your highest bonus earned: %d\n", user.getHighestBonus())

	//---------------------------------
}
