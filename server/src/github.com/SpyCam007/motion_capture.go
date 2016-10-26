package SpyCam007

import (
	"net/http"
	"fmt"
	"errors"
)

func HandlePicture(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprint(w, "Got to save picture via a post request.", r.Method)
	} else if r.Method == "GET" {
		fmt.Fprint(w, "Got to save picture via a get request. ", r.Method)
	} else {
		fmt.Fprint(w, "Got to save picture with an invalid request type ", r.Method)
	}
}

func savePicture(r *http.Request) error {
	if (r.Method != "POST") {
		return errors.New("The given request for savePicture was not a post request.")
	}
	if (r.Body == nil) {
		return errors.New("The given request for savePicture has a nil body. Picture expected.")
	}


	return nil
}
