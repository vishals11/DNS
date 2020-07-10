# From base image
FROM golang:1.12 AS builder
#setting up working directory

WORKDIR /go/src/DroneNavigationSystem
ADD . /go/src/DroneNavigationSystem

RUN go get -d -v
RUN go build -o server
RUN cp server /server

EXPOSE 8000
ENTRYPOINT ["/server"]