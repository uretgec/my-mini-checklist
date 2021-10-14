FROM golang:1.17-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /go/src/myminichecklist/db

WORKDIR /go/src/myminichecklist

COPY . .

RUN go install -v ./...

FROM alpine:latest

RUN mkdir -p /usr/local/my-mini-checklist

COPY --from=builder /go/bin/my-mini-checklist /usr/local/my-mini-checklist/my-mini-checklist

WORKDIR /usr/local/my-mini-checklist

ENTRYPOINT [ "./my-mini-checklist",  "--addr", ":3000", "--dbpath", "./db/store.db", "--bgsave", "1m0s", "--logpath", "./db/service-http.log" ]