# uber-eats-like
golang学習のためのウーバーイーツ風アプリ

## mysqlに入る

```
$ docker compose exec db mysql -u<ユーザー名> -p<パスワード> <データベース名>
```

## マイグレーションツール

golang-migrateを使用

マイグレーションファイル作成コマンド

```
$ migrate create -ext sql -dir db/migrations -seq create_hoge_table
```

マイグレーションの適用

```
$ migrate -path db/migrations -database "mysql://myuser:mypassword@tcp(db:3306)/mydb" up
```

マイグレーションのロールバック

```
$ migrate -path db/migrations -database "mysql://myuser:mypassword@tcp(db:3306)/mydb" down
```

