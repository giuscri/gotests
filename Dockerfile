FROM golang:alpine

WORKDIR /app
COPY . .

RUN go test -v ./... -bench=. -cover
