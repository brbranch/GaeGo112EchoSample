# GAE/Go1.12+Echo
## Description
GAE+Go1.12でEchoサーバーを立ち上げるためのサンプルです。
DataStoreへのアクセスはmercari/datastoreのboomを利用しています。

## Requirements
* goenv 2.0.0beta11
* go v1.12.9
* Google Cloud SDK 253.0.0

## ローカル環境での確認
### Datastore Emulatorの起動
```
gcloud beta emulators datastore start --host-port localhost:8059 --project test-project
```

### ローカルサーバーの起動

```
$ cd ./src
$ env DATASTORE_EMULATOR_HOST=localhost:8059 DATASTORE_PROJECT_ID=test-project go run main.go
```

## deploy
app.yamlにある `PROJECT_ID` をご自身のGCP Project IDに変更下さい。

```
gcloud app deploy
```

### バージョン指定してでのデプロイ(Trafficを移行させない)

```
gcloud app deploy --project PROJECT_ID --version VERSION_ID --no-promote
```