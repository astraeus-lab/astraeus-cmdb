FROM golang:1.21.5 AS builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0

FROM alpine:3.19.0
