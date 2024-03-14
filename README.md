# sandbox-go-api


## Commands

start

`go run main.go`

MySQLを起動

`docker-compose up -d mysql`

MySQLを操作

`docker-compose run mysql-cli`

### Migration

golang-migrateのインストール(Mac)

`brew install golang-migrate`

マイグレーションファイル作成

`migrate create -ext sql -dir database/migrations -seq [Migration File Name]`

up

`migrate --path database/migrations --database 'mysql://sandbox:develop@tcp(127.0.0.1:3306)/sandbox-go-api' -verbose up`

down

`migrate --path database/migrations --database 'mysql://sandbox:develop@tcp(127.0.0.1:3306)/sandbox-go-api' -verbose down`
