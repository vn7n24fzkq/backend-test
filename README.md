# backend-test

## How to run

直接執行會使用 sqlite 當作資料庫
因為直接將 web 呼叫 API 的位置設定為 127.0.0.1:8080，所以啟動之後會使用 8080 port

```sh
go run main.go
```

並且因為開發時是使用 MariaDB ， 所以如果有 MariaDB 的環境也可以透過加入 dsn 來使用 MariaDB

ie. dsn={username}:{password}@tcp({host}:{port})/{database}?charset=utf8mb4&parseTime=True&loc=Local

```sh
dsn='admin:password@tcp(127.0.0.1:3306)/backend_test?charset=utf8mb4&parseTime=True&loc=Local' go run main.go
```
