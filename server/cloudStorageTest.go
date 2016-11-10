// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//[START sample]
// Package gcsdemo is an example App Engine app using the Google Cloud Storage API.
//
// NOTE: the cloud.google.com/go/storage package is not compatible with
// dev_appserver.py, so this example will not work in a local development
// environment.
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
)


// bucket is a local cache of the app's default bucket name.
//var bucket string // or: var bucket = "<your-app-id>.appspot.com"

// demo struct holds information needed to run the various demo functions.
type CSHandler struct {
  bucketName string
	bucket *storage.BucketHandle
	client *storage.Client
	w   io.Writer
  	b   *bytes.Buffer
	ctx context.Context
	// cleanUp is a list of filenames that need cleaning up at the end of the demo.
	cleanUp []string
	// failed indicates that one or more of the demo steps failed.
	failed bool
}

func (d *CSHandler) errorf(format string, args ...interface{}) {
	d.failed = true
	fmt.Fprintln(d.w, fmt.Sprintf(format, args...))
	log.Errorf(d.ctx, format, args...)
}


func (d * CSHandler)InitCSHandler(r *http.Request, bucketName string) error {
  //[START get_default_bucket]
	ctx := appengine.NewContext(r)
	if bucketName == "" {
		var err error
		if bucketName, err = file.DefaultBucketName(ctx); err != nil {
			log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
			return err
		}
	}
	//[END get_default_bucket]

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to create client: %v", err)
		return err
	}
  buf := &bytes.Buffer{}
  d.bucketName = bucketName
	d.w = buf
  d.b = buf
  d.ctx = ctx
  d.client = client
  d.bucket =  client.Bucket(bucketName)
  return nil
}

func (d *CSHandler) Close() {
  if d.client != nil {
      d.client.Close()
  }
}
// handler is the main demo entry point that calls the GCS operations.
func Handler(w http.ResponseWriter, r *http.Request) {
	if appengine.IsDevAppServer() {
		http.Error(w, "This example does not work with dev_appserver.py", http.StatusNotImplemented)
	}


	//defer client.Close()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//fmt.Fprintf(w, "Demo GCS Application running from Version: %v\n", appengine.VersionID(ctx))
	//fmt.Fprintf(w, "Using bucket name: %v\n\n", bucket)
 d := &CSHandler{}
 d.InitCSHandler(r, "")
 defer d.Close()
	n := "demo-testfile-go"
	d.CreateCSFile(n,   []byte("Lol this is a fake file."), "text/plain")
	d.deleteFile(n)
	if d.failed {
		w.WriteHeader(http.StatusInternalServerError)
		d.b.WriteTo(w)
		fmt.Fprintf(w, "\nDemo failed.\n")
	} else {
		w.WriteHeader(http.StatusOK)
		d.b.WriteTo(w)
		fmt.Fprintf(w, "\nDemo succeeded.\n")
	}
}

//[START write]
// createFile creates a file in Google Cloud Storage.
func (d *CSHandler) CreateCSFile(fileName string, fileContents []byte, contentType string) {
	fmt.Fprintf(d.w, "Creating file /%v/%v\n", d.bucketName, fileName)
	wc := d.bucket.Object(fileName).NewWriter(d.ctx)
	wc.ContentType = contentType
	if _, err := wc.Write(fileContents); err != nil {
		d.errorf("createFile: unable to write data to bucket %q, file %q: %v", d.bucketName, fileName, err)
		return
	}
  if err := wc.Close(); err != nil {
    	d.errorf("error closing writer  %v", err)
  }
}

//[END write]

// deleteFiles deletes all the temporary files from a bucket created by this CSHandler.
func (d *CSHandler) deleteFile(fileName string) {
	io.WriteString(d.w, "\nDeleting files...\n")
	fmt.Fprintf(d.w, "Deleting file %v\n", fileName)
	if err := d.bucket.Object(fileName).Delete(d.ctx); err != nil {
		d.errorf("deleteFiles: unable to delete bucket %q, file %q: %v", d.bucketName, fileName, err)
		return
	}
}

//[END sample]
