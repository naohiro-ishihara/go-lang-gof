FROM golang:1.20.4-alpine

# アップデートとgitのインストール
RUN apk --update-cache --no-cache add git

# set alias commands
ENV ENV='/etc/profile'
COPY .aliases .
RUN cat .aliases >> 'ENV'

RUN mkdir -p /work

WORKDIR /work
