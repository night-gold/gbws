package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["message"]
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	//http.HandleFunc("/", sayHello)
	r := newRouter()
	http.ListenAndServe(":8080", r)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	//	panic(err)
	//}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", sayHello).Methods("GET")
	r.HandleFunc("/hello/{message}", sayHello).Methods("GET")
	r.HandleFunc("/modele", templating).Methods("GET")
	return r
}

func templating(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}
