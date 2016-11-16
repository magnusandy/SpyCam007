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

/** Return a list of pictures matching the query. If err != nil, something went wrong when querying. */
func QueryForPictureByDate(begin time.Time, end time.Time, r *http.Request) ([]SpycamPicture, error) {
 ctx := appengine.NewContext(r)
 var pictures []SpycamPicture
 q := datastore.NewQuery(SpycamModelName).
 Filter("Timestamp <=", end).
 Filter("Timestamp >=", begin).
 Order("Timestamp") // Order("-Timestamp")
 _, err := q.GetAll(ctx, &pictures) // returns keys, and err.
 return pictures, err
}
