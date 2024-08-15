# Starter

###### cmd
``` shell
git clone https://github.com/Maxl3oss/go-gin-swagger-template.git
cd ./go-gin-swagger-template.git
go mod tidy
```

# Dot env file
###### create .env file

```yaml
Port=[port gogin]
Connect_type=psql
Db_psql=Postgres://[Username]:[Password]@[Hostname]:[Port]/[Database_Name]?Sslmode=Disable

POSTGRES_USER=[postgres]
POSTGRES_PASSWORD=[password]
POSTGRES_DB=[Database_Name]
```

# Gen swagger
###### install gin-swagger
https://github.com/swaggo/gin-swagger
> after that  run
```shell
swag init -g ./cmd/api/main.go -o ./docs
```

###### install go air for live reload
https://github.com/air-verse/air
>after that run
```shell
air init
```

###### update file .air.toml
> after that run
```yaml
cmd = "go build -o ./tmp/main ./cmd/api"
```
```shell
docker compose up -d
air
```
