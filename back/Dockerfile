FROM golang:1.22.1

WORKDIR /go/src

COPY go* ./

RUN go mod download
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm

COPY . .