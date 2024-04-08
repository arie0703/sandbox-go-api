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

`migrate --path database/migrations --database 'mysql://main:develop@tcp(127.0.0.1:3306)/sandbox-go-api' -verbose up`


down

`migrate --path database/migrations --database 'mysql://main:develop@tcp(127.0.0.1:3306)/sandbox-go-api' -verbose down`

### Replication

コンテナのIPアドレス確認

`docker container exec -it sandbox-go-api_mysql-main_1 hostname -i`

Main環境

```sql
# レプリケーション用ユーザー作成
CREATE USER 'replica'@'172.30.%.%' IDENTIFIED BY 'replica';
GRANT REPLICATION SLAVE ON *.* TO 'replica'@'172.30.%.%';
GRANT ALL PRIVILEGES ON *.* TO 'replica'@'172.30.%.%';
# バイナリログの情報確認
SHOW MASTER STATUS;

# Authentication requires secure connection. というエラーが出た時の対処法
ALTER USER 'replica'@'172.30.%.%' IDENTIFIED WITH mysql_native_password BY 'password';
```

Replica環境

```sql
CHANGE MASTER TO MASTER_HOST='mysql-main', MASTER_PORT=3306, MASTER_LOG_FILE='5f7556518596-bin.000001', MASTER_LOG_POS=156;

START SLAVE USER = 'replica' PASSWORD = 'password';

STOP SLAVE;

SHOW SLAVE STATUS\G;

# レプリケーション設定のリセット
RESET SLAVE;
```

