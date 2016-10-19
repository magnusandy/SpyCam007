package SpyCam007

import (
	"net/http"
	"fmt"
)

func SavePicture(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprint(w, "Got to save picture via a post request.", r.Method)
	} else if r.Method == "GET" {
		fmt.Fprint(w, "Got to save picture via a get request. ", r.Method)
	} else {
		fmt.Fprint(w, "Got to save picture with an invalid request type ", r.Method)
	}
}
