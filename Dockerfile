FROM golang:1.7.3
ADD . /go/src/blog-service
WORKDIR /go/src/blog-service
RUN go get
RUN go install blog-service
ENTRYPOINT /go/bin/blog-service
EXPOSE 8080
