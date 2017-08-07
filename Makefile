NAME=mikedotcom-blog
VERSION=v1.0.0
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
