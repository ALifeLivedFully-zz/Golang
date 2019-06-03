package main

import (
	"fmt"
	"net/http"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, go is neat!")
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the about page!")
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/about/", abouthandler)
	http.ListenAndServe(":8000", nil)
}