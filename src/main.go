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
	"fmt"
	"gaego112echosample/client"
	"gaego112echosample/handler"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
)

func main() {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")

	client.InitRedis(redisAddr, redisPass)

	e := echo.New()
	http.Handle("/", e)

	e.GET("/", handler.HelloWorld)

	// Redisの操作
	e.GET("/redis/put/:name", func(e echo.Context) error {
		name := e.Param("name")
		if err := client.GetRedisClient().PutString("test", name); err != nil {
			log.Fatalf("faield to get redis (reason: %v)", err)
			return e.String(http.StatusInternalServerError, "error")
		}
		return e.String(http.StatusOK, fmt.Sprintf("put redis: %s", name))
	})
	e.GET("/redis/get", func(e echo.Context) error {
		if name, err := client.GetRedisClient().GetString("test"); err != nil {
			log.Fatalf("faield to get redis (reason: %v)", err)
			return e.String(http.StatusInternalServerError, "error")
		} else {
			return e.String(http.StatusOK, fmt.Sprintf("get redis: %s", name))
		}
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
