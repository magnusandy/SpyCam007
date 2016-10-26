package main

import (
	"fmt"
	"github.com/SpyCam007"
	"net/http"
)

func init() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/pictures/", SpyCam007.HandlePicture)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func main() {

}
