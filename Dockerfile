FROM golang:1.7.3
MAINTAINER Mike Cutalo (mcutalo88@gmail.com)

ADD . /go/src/blog-service
WORKDIR /go/src/blog-service

RUN curl https://glide.sh/get | sh
RUN glide install
RUN go build -o /go/bin/blog-service

ENTRYPOINT /go/bin/blog-service
EXPOSE 8080
