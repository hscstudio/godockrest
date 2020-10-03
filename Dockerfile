FROM golang:alpine
RUN apk add git

ENV GOBIN /go/bin

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/labstack/echo/...
