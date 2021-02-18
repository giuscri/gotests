FROM golang:alpine

RUN apk add git && go get -u github.com/kisielk/errcheck

WORKDIR /app

COPY . .

RUN errcheck . && go test -v ./... -bench=. -cover
