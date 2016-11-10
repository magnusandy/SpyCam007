package main

import (
  "google.golang.org/appengine/datastore"
  "google.golang.org/appengine"
	"time"
	"fmt"
  "net/http"
)

const (
	SpycamModelName = "SpycamPicture"
)

type SpycamPicture struct {
	Timestamp time.Time		// When the picture was received.
	Url 	  string		// Where the files are stored.
}

func (pictureToSave *SpycamPicture) Save(r *http.Request) {
	// TODO: Do I need to fill in information?
	ctx := appengine.NewContext(r)


	key := datastore.NewIncompleteKey(ctx, SpycamModelName, nil)

	if _, err := datastore.Put(ctx, key, pictureToSave); err != nil {
		//TODO logging for error.
		fmt.Printf("Error when saving model: %s\n", err.Error())
	} else {
		fmt.Println("Successful.")
	}
}
