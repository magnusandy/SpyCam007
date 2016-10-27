package main

import (

	"github.com/SpyCam007"
	"net/http"
  "log"
  "golang.org/x/net/context"
  "golang.org/x/oauth2/google"
  storage "google.golang.org/api/storage/v1"
)


func init() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/pictures/", SpyCam007.SavePicture)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
