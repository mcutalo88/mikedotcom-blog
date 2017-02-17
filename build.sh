#!/bin/bash

cd "$(dirname "$0")"

go get
go build -o bin/blog-service
