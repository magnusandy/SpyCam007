package main

import (
	"net/http"
	"fmt"
	"errors"
	"io/ioutil"
  "google.golang.org/appengine"
)

func HandlePicture(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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

	imageFile, header, err := r.FormFile("picture") // img is the key of the form-data
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return nil
	}
  fmt.Fprintf(w, "Got filename %s\n", header.Filename)
  if appengine.IsDevAppServer() {
    fmt.Fprintf(w, "Error: %s\n", "Cannot save picture on dev app server")
	}else{
    csHandler := CSHandler{}
    csHandler.InitCSHandler(r, "")
    content, err := ioutil.ReadAll(imageFile)
    if err != nil {
      fmt.Fprintf(w, "Error ReadAll image contents: %s\n", "Cannot save")
      return nil
    }
    csHandler.CreateCSFile(header.Filename, content, "image/png")
  }
	return nil
}
