FROM golang:1.13-alpine

WORKDIR /go/src/app

COPY . .

RUN go build -o /service ./cmd/service

EXPOSE 8090

ENTRYPOINT ["/service"]