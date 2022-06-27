package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                     string
	Age                      uint16
	Money                    int16
	AverageGrades, Happiness float32
	Hobbies                  []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money "+
		"equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	val := User{"Valentin", 25, 19876, 4.2, 0.8, []string{"Wath Movie", "Learn Music", "Football"}}
	// val.setNewName("Valyok")
	// fmt.Fprintf(w, `<h1>Main Text</h1>
	// <b>Main Text</b>`)
	tmpl, _ := template.ParseFiles("templates/homePage2.html")
	tmpl.Execute(w, val)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":1488", nil)
}

func main() {
	// val := User{name: "Valentin", age: 25, money: 19876, averageGrades: 4.2, happiness: 0.8}

	handleRequest()
}
