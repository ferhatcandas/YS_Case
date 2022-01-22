#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /
COPY ./src .
RUN go build -o main

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder / /
ENTRYPOINT /main
LABEL Name=yscase Version=0.0.1
EXPOSE 3000


