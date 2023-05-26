FROM golang:1.20.4-alpine

# apk uddate and install tools
RUN apk --update-cache --no-cache add git

# envroiments
ENV ENV="/etc/profile"

# set alias commands
COPY .aliases .
RUN cat .aliases >> "$ENV"

# set golang binary path
RUN echo 'export PATH="$PATH:/usr/local/go/bin"' >> "$ENV"

RUN mkdir -p /work

WORKDIR /work
