package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
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
	vars := chi.URLParam(r, "message")
	message := vars
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

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.HandleFunc("/", sayHello)
	r.HandleFunc("/hello/{message}", sayHello)
	r.HandleFunc("/modele", templating)
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))
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
