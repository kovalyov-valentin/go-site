package main 

import (
	"fmt"
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func handleFunc() {
	//Функция отслеживания страницы
	http.HandleFunc("/", index)
	//Данный метод позволяет нам вызвать определенный сервер
	http.ListenAndServe(":8080", nil)
}

func main (
	handleFunc()
)