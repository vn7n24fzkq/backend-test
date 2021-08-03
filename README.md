# backend-test

- 網頁位址 : http://127.0.0.1:8080/static/

## Requirements

- go 1.16

### Optional
- Mariadb 10.3.3

## How to run

為了方便執行，直接執行會使用 sqlite 當作資料庫
因為直接將 web 呼叫 API 的位置設定為 127.0.0.1:8080，所以這邊固定使用 8080 port

```sh
go run main.go
```

並且開發時是使用 MariaDB ， 所以如果有 MariaDB 的環境也可以透過加入 dsn 來選擇使用 MariaDB

e.g. dsn={username}:{password}@tcp({host}:{port})/{database}?charset=utf8mb4&parseTime=True&loc=Local

```sh
dsn='admin:password@tcp(127.0.0.1:3306)/backend_test?charset=utf8mb4&parseTime=True&loc=Local' go run main.go
```

## How to test

```sh
go test ./...
```
