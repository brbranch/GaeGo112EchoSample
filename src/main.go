// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"go.mercari.io/datastore/boom"
	"go.mercari.io/datastore/clouddatastore"
	"log"
	"net/http"
	"os"
)

type Post struct {
	Kind	string `datastore:"-" boom:"kind,post"`
	ID		int64 `datastore:"-" boom:"id"`
	Content	string
}

func main() {
	e := echo.New()
	http.Handle("/", e)

	projectId := os.Getenv("PROJECT_ID")
	if projectId == "" {
		projectId = "jankenonline"
	}

	e.GET("/", func (e echo.Context) error {
		ctx := context.Background()
		dataClient, err := datastore.NewClient(ctx, projectId)
		if err != nil {
			log.Fatalf("failed to get client (reason: %v)", err)
			return e.String(http.StatusInternalServerError, "error")
		}
		client, err := clouddatastore.FromClient(ctx, dataClient)
		if err != nil {
			log.Fatalf("failed to get datastoreclient (reason: %v)", err)
			return e.String(http.StatusInternalServerError, "error")
		}
		defer client.Close()
		b := boom.FromClient(ctx, client)
		post := &Post{Content:"test"}
		if _, err := b.Put(post); err != nil {
			log.Fatalf("failed to put datastore (reason: %v)", err)
			return e.String(http.StatusInternalServerError, "error")
		}

		return e.String(http.StatusOK, "Hello Echo!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}
