FROM golang:1.12.7-alpine3.10

WORKDIR /go/src

RUN apk add --no-cache \
        alpine-sdk \
        git

ENV GO111MODULE=on \
        GOOGLE_APPLICATION_CREDENTIALS=/go/src/batch/sql-service-account.json
COPY batch /go/src/batch
