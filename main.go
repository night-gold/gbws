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
	r := newRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	//Simple sayHello function on GET method
	r.HandleFunc("/", sayHello).Methods("GET")
	//Say Hello and add the message part of URI
	r.HandleFunc("/hello/{message}", sayHello).Methods("GET")
	//show the template layout
	r.HandleFunc("/modele", templating).Methods("GET")
	//open a link for serving static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))
	//create a form template
	//r.HandleFunc("/form", formTest)
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

//func formTest(w http.ResponseWriter, r *http.Request) {
// tmpl := template.Must(template.ParseFiles("form.html"))
//}
