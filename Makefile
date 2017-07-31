NAME=mikedotcom-blog
VERSION=v0.0.2
DOCKER_IMAGE=registry.azeroth.io/botnetz/$(NAME):$(VERSION)

.PHONY: install build run image push release

install:
	glide install
	glide up -v

build:
	go build -o bin/$(NAME)

run:
	./bin/$(NAME)

image:
	docker build -t $(DOCKER_IMAGE) .

push:
	docker push $(DOCKER_IMAGE)

release: image push
