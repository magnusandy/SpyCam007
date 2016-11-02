package main

import (
	"net/http"
	"fmt"
	"errors"
	//"time"
)

func HandlePicture(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprint(w, "Got to save picture via a post request.", r.Method)
		savePicture(r, w)
	} else if r.Method == "GET" {
		fmt.Fprint(w, "Got to save picture via a get request. ", r.Method)
	} else {
		fmt.Fprint(w, "Got to save picture with an invalid request type ", r.Method)
	}

}

func savePicture(r *http.Request, w http.ResponseWriter) error {
	if (r.Method != "POST") {
		return errors.New("The given request for savePicture was not a post request.")
	}

	//file, handler, err := r.FormFile("img") // img is the key of the form-data
	_, handler, err := r.FormFile("picture") // img is the key of the form-data
	if err != nil {
		// TODO add logging.
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return nil
	}
	fmt.Println("File is good")
	fmt.Println(handler.Filename)
	fmt.Println()
	fmt.Println(handler.Header)
	fmt.Fprintf(w, "Got filename %s\n", handler.Filename)
	return nil
}
