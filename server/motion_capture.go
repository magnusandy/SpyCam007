package main

import (
	"net/http"
	"fmt"
	"errors"
	"io/ioutil"
  "google.golang.org/appengine"
  "time"
  "encoding/json"
  "strconv"
)

func HandlePicture(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		savePicture(r, w)
	} else if r.Method == "GET" {
		query := r.URL.Query()
    startTime := query.Get("startTime")
    endTime := query.Get("endTime")
    if startTime == "" || endTime == "" {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    }else{
      startTimestamp, errStart := strconv.ParseInt(startTime, 10, 64)
      endTimestamp, errEnd := strconv.ParseInt(endTime, 10, 64)
      if errStart != nil || errEnd != nil {
        http.Error(w, "failed to convert timestamps", http.StatusBadRequest)
      }
      pictures, err := QueryForPictureByDate(time.Unix(startTimestamp, 0), time.Unix(endTimestamp, 0), r)
      if err != nil{
        http.Error(w, "failed to get query of pictures", http.StatusInternalServerError)
      }

      jsonResp, err := json.Marshal(pictures)
      if err != nil {
        http.Error(w, "Failed to marshal json response", http.StatusInternalServerError)
      }

      w.Header().Set("Content-Type", "application/json")
      w.Write(jsonResp)
    }

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
    mediaLink, err := csHandler.CreateCSFile(header.Filename, content, "image/png")
    if err != nil {
      fmt.Fprintf(w, "Error saving the image to storage: %s\n", err.Error())
    }

    newPicModel := &SpycamPicture{Url:mediaLink, Timestamp: time.Now().UTC()}
    fmt.Fprintf(w, "newImageURL: %s\n", newPicModel.Url)
    newPicModel.Save(r)
  }
	return nil
}
