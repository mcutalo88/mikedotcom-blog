NAME=mikedotcom-blog
# VERSION := $(shell git describe --abbrev=0 HEAD)
VERSION := $(shell git name-rev --tags --name-only $(shell git rev-parse HEAD))

# DOCKER_IMAGE=registry.azeroth.io/botnetz/$(NAME):$(VERSION)
DOCKER_IMAGE=mcutalo/mikedotcom-blog:$(VERSION)

.PHONY: install build run image push release

install:
	glide install
	glide up -v

build:
	go build -o bin/$(NAME)

run:
	export GO_ENV=dev
	./bin/$(NAME)

image:
	docker build -t $(DOCKER_IMAGE) .

push:
	docker push $(DOCKER_IMAGE)

release: image push
