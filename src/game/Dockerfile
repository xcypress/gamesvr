FROM registry.cn-hangzhou.aliyuncs.com/codoon/docker-golang:latest
MAINTAINER xinhp <xinhp.git@gmail.com>
COPY . /go/src/game
RUN go install game
ENTRYPOINT /go/bin/game
