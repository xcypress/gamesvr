FROM golang:latest
MAINTAINER xinhp <xinhp.git@gmail.com>
COPY . /go/src/game
RUN go install game

