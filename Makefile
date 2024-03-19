APP_NAME=logger
GOLANG_VERSION=latest

build:
	@echo ":::building image"
	docker build . --build-arg APP_NAME=$(APP_NAME) --build-arg GOLANG_VERSION=$(GOLANG_VERSION)

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*
