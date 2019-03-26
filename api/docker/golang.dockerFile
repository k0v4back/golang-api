FROM golang:latest

RUN apt-get update
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"