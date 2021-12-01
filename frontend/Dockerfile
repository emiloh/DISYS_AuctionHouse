FROM golang:1.17-alpine AS base

WORKDIR /src

COPY go.mod /src/
COPY go.sum /src/

RUN go mod download

COPY . /src/

RUN go build -o "bin/out" github.com/emiloh/DISYS_AuctionHouse/frontend
EXPOSE 1400
ENTRYPOINT ["bin/out"]