SERVICE_NAME=teamtrack
CWD:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: rebuild

kit-clear:
	rm -f cmd/service/service_gen.go \
		pkg/endpoint/endpoint.go \
		pkg/endpoint/endpoint_gen.go \
		pkg/endpoint/middleware.go \
		pkg/http/handler.go \
		pkg/http/handler_gen.go \
		pkg/service/middleware.go

.ONESHELL:
kit-gen:
	cd "$(CWD)/.."
	kit generate service teamtrack --dmw
	cd -

dep:
	dep ensure -v

gen:
	go generate ./...

build:
	rm -f out/*
	env GOOS=linux GOARCH=amd64 go build -o out/${SERVICE_NAME} cmd/main.go

rebuild: kit-clear kit-gen dep gen build

install:
	cp ./out/${SERVICE_NAME} /usr/local/bin

start:
	docker-compose up --build