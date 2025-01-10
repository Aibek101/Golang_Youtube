package main

import (
	"fmt"
	"net/http"
)

func main() {
	// var bob User = {name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}
	// bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}

	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	bob.setNewName("Alex")
	fmt.Fprintf(w, bob.getAllInfo())
}
func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "President - +7(777)777-77-77")
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d. He has %d$.", u.name, u.age, u.money)
}
func (u *User) setNewName(newName string) {
	u.name = newName
}

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}
