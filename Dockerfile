FROM golang:alpine

WORKDIR /go/src
COPY . /go/src/discord.land/site

RUN apk update && apk upgrade && apk add git --no-cache
RUN go get -d discord.land/site/server
RUN go install discord.land/site/server

ENTRYPOINT /go/bin/server

EXPOSE 8002