#!/bin/bash

cd "$(dirname "$0")"

go get -insecure
go build -o bin/blog-service

docker build -t registry.azeroth.io/azeroth/blog-service:v1.0.0 .
docker push registry.azeroth.io/azeroth/blog-service:v1.0.0
