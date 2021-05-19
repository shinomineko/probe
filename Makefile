SHELL := bash

.PHONY: all
all: build push

.PHONY: build
build:
	docker build -t shinomineko/probe:latest -t quay.io/shinomineko/probe:latest .

.PHONY: push
push:
	docker push -a shinomineko/probe
	docker push -a quay.io/shinomineko/probe
