package main

import (
	"fmt"
	"net/http"
)

func indexhandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,`
	<h1>heres a heading</h1>
	<p>this is a <strong>paragraph</strong></p>
	<p>this paragraph has a %s %s. cool huh?</p>
	`,"few","variables")
}

func main()  {
	http.HandleFunc("/", indexhandler)
	http.ListenAndServe(":8000", nil)
}