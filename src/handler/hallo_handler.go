package handler

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	client2 "gaego112echosample/client"
	"github.com/labstack/echo"
	datastore2 "go.mercari.io/datastore"
	"go.mercari.io/datastore/boom"
	"go.mercari.io/datastore/clouddatastore"
	"go.mercari.io/datastore/dsmiddleware/rediscache"


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

	// Redisとの連携を行う
	redisConn := client2.GetRedisClient().GetConnection()
	defer redisConn.Close()
	mw := rediscache.New(redisConn,
		// ログが出るようにする
		rediscache.WithLogger(func(ctx context.Context, format string, args ...interface{}){
			log.Printf(format, args...)
		}),
		// Redisに登録されてるか見るためにログに出力するようにしてる
		rediscache.WithCacheKey(func(key datastore2.Key) string {
			cacheKey := fmt.Sprintf("cache:%s", key.Encode())
			log.Printf("redis cache key: %s", cacheKey)
			return cacheKey
		}),
	)
	client.AppendMiddleware(mw)

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

func HelloRedis(e echo.Context) error {


	return nil
}