ARG ARCH="amd64"
ARG OS="linux"

FROM golang:1.17 AS build

WORKDIR /httpserver
COPY . .

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

RUN GOOS=linux go build -installsuffix cgo -o /httpserver/httpserver /httpserver/cmd/http_server/main.go

FROM centos:centos8.2.2004

ENV LANG=en_US.UTF-8

LABEL maintainer="Shijjie Zhang <1550146843@qq.com>"

COPY --from=build /httpserver/httpserver /httpserver/httpserver
COPY --from=build /httpserver/conf/config.yaml /httpserver/config.yaml

WORKDIR /httpserver/

ENTRYPOINT ["/httpserver/httpserver", "--config=/httpserver/config.yaml"]

EXPOSE 8088