﻿# httpserver

## start
```shell
httpserver --config={{config_file}}
```

## docker build
```shell
docker build -t httpserver:v1 -f .\Dockerfile .\ 
```

## go build
```shell
go build -installsuffix cgo -o /httpserver/httpserver /httpserver/cmd/http_server/main.go
```

## api
```text
GET    /api/v1/zsj/healthz
GET    /api/v1/zsj/data
POST   /api/v1/zsj/data
```
