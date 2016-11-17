package main


import (
	"net/http"
)

func init() {
  http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", helloWorld)
  http.HandleFunc("/testing", Handler)
	http.HandleFunc("/pictures/", HandlePicture)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
