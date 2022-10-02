# next-gin-mysql

## For Development
Place the `.env` file, one under `backend`, another under `migration`

```
HTTP_HOST=""
HTTP_PORT=8080
DB_HOST="mysql"
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=root
DB_DATABASE=backend
```
## air上手くいかなかった，ので消した．
## コンテナ起動
## docker-compose up --build -d
## コンテナの中に入る
## docker exec -it next-gin-mysql-go-1 /bin/bash
## backend直下のmainを動かす
## go run main.go
## 同じコンテナに別のターミナルから入る
## docker exec -it next-gin-mysql-go-1 /bin/bash
## migration直下のmainを動かす
## go run main.go
## postmanで試す