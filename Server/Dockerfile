FROM golang:1.17-alpine AS base

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o "bin/out" github.com/emiloh/DISYS_AuctionHous/tree/simpler/Server
EXPOSE 8000
ENTRYPOINT ["bin/out"]