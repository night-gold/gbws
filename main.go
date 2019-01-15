package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

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
	r.HandleFunc("/{message}", sayHello).Methods("GET")
	return r
}
