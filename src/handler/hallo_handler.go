package handler

import (
	"cloud.google.com/go/datastore"
	"github.com/labstack/echo"
	"go.mercari.io/datastore/boom"
	"go.mercari.io/datastore/clouddatastore"
	"log"
	"net/http"
	"os"
)

type Post struct {
	Kind	string `datastore:"-" boom:"kind,post" json:"-"`
	ID		int64 `datastore:"-" boom:"id" json:"id"`
	Content	string `datastore:"content" json:"content"`
}

func HelloWorld(e echo.Context) error {
	// app.yaml などに設定しておく
	projectId := os.Getenv("PROJECT_ID")
	log.Printf("Project ID: %s", projectId)
	// DataStore Clientの作成
	ctx := e.Request().Context()
	dataClient, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("failed to get client (reason: %v)", err)
		return e.String(http.StatusInternalServerError, "error")
	}
	// mercari.datastoreでラップする
	client, err := clouddatastore.FromClient(ctx, dataClient)
	if err != nil {
		log.Fatalf("failed to get datastoreclient (reason: %v)", err)
		return e.String(http.StatusInternalServerError, "error")
	}
	defer client.Close()
	// boomを利用
	b := boom.FromClient(ctx, client)
	post := &Post{ID: 12345, Content:"test"}
	// 保存
	if _, err := b.Put(post); err != nil {
		log.Fatalf("failed to put datastore (reason: %v)", err)
		return e.String(http.StatusInternalServerError, "error")
	}
	// 取得
	getPost := &Post{ID: 12345}
	if err := b.Get(getPost); err != nil {
		log.Fatalf("failed to get datastore (reason: %v)", err)
		return e.String(http.StatusInternalServerError, "error")
	}

	return e.JSON(http.StatusOK, getPost)
}
