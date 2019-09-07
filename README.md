# GAE/Go1.12+Echo
## Description
GAE+Go1.12でEchoサーバーを立ち上げるためのサンプルです。
DataStoreへのアクセスはmercari/datastoreのboomを利用しています。

### Qiita
#### gae_echo
https://qiita.com/br_branch/items/a26480a05ecb97ac20b3
#### feature/redis


## Requirements
* goenv 2.0.0beta11
* go v1.12.9
* Google Cloud SDK 253.0.0
* Redis local server 5.0.5

### Redisのインストール
see https://redis.io/topics/quickstart

## ローカル環境での確認

ターミナル3つ使ってそれぞれのエミュレーター / サーバーを立ち上げます。

### Redis Serverの起動
```
$ redis-server
65146:C 07 Sep 2019 17:03:08.988 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
65146:C 07 Sep 2019 17:03:08.988 # Redis version=5.0.5, bits=64, commit=00000000, modified=0, pid=65146, just started
65146:C 07 Sep 2019 17:03:08.988 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
65146:M 07 Sep 2019 17:03:08.989 * Increased maximum number of open files to 10032 (it was originally set to 4864).
                _._
           _.-``__ ''-._
      _.-``    `.  `_.  ''-._           Redis 5.0.5 (00000000/0) 64 bit
  .-`` .-```.  ```\/    _.,_ ''-._
 (    '      ,       .-`  | `,    )     Running in standalone mode
 |`-._`-...-` __...-.``-._|'` _.-'|     Port: 6379
 |    `-._   `._    /     _.-'    |     PID: 65146
  `-._    `-._  `-./  _.-'    _.-'
 |`-._`-._    `-.__.-'    _.-'_.-'|
 |    `-._`-._        _.-'_.-'    |           http://redis.io
  `-._    `-._`-.__.-'_.-'    _.-'
 |`-._`-._    `-.__.-'    _.-'_.-'|
 |    `-._`-._        _.-'_.-'    |
  `-._    `-._`-.__.-'_.-'    _.-'
      `-._    `-.__.-'    _.-'
          `-._        _.-'
              `-.__.-'

65146:M 07 Sep 2019 17:03:08.990 # Server initialized
65146:M 07 Sep 2019 17:03:08.990 * Ready to accept connections
^C65146:signal-handler (1567843408) Received SIGINT scheduling shutdown...
65146:M 07 Sep 2019 17:03:28.793 # User requested shutdown...
65146:M 07 Sep 2019 17:03:28.793 * Saving the final RDB snapshot before exiting. 
```

### Datastore Emulatorの起動
```
$ gcloud beta emulators datastore start --host-port localhost:8059 --project test-project
```

### ローカルサーバーの起動

```
$ cd ./src
$ env DATASTORE_EMULATOR_HOST=localhost:8059 DATASTORE_PROJECT_ID=test-project REDIS_ADDR=localhost:6379 go run main.go
```

## deploy
app.yamlにある環境変数をそれぞれ設定ください。

* `PROJECT_ID` : ご自身のGCP Project ID
* `REDIS_ADDR` : Redisのエンドポイント
* `REDIS_PASS` : Redisのパスワード

```
gcloud app deploy
```

### バージョン指定してでのデプロイ(Trafficを移行させない)

```
gcloud app deploy --project PROJECT_ID --version VERSION_ID --no-promote
```
