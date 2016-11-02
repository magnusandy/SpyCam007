package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", helloWorld)
  http.HandleFunc("/testing", Handler)
	http.HandleFunc("/pictures/", HandlePicture)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
