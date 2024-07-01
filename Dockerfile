FROM golang:1.19.1
FROM ubuntu:latest

WORKDIR app
COPY ./transloapi ./
EXPOSE $PORT
ENTRYPOINT ["./transloapi"]