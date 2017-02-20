#!/bin/bash

cd "$(dirname "$0")"

go get -insecure
go build -o bin/blog-service
