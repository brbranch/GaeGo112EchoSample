# GAE/Go1.12+Echo
## Description
GAE+Go1.12でEchoサーバーを立ち上げるためのサンプルです。
DataStoreへのアクセスはmercari/datastoreのboomを利用しています。

## Requirements
* goenv 2.0.0beta11
* go v1.12.9
* Google Cloud SDK 253.0.0

## ローカル環境での確認
```
$ cd ./src
$ go run main.go
```

## deploy

```
gcloud app deploy
```

### バージョン指定してでのデプロイ(Trafficを移行させない)

```
gcloud app deploy --project PROJECT_ID --version VERSION_ID --no-promote
```