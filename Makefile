SERVICE_NAME=teamtrack

all: build

dep:
	dep ensure -v

gen:
	go generate ./...

build: dep gen
	rm -f out/
	env GOOS=linux GOARCH=amd64 go build -o out/${SERVICE_NAME} cmd/main.go

install:
	cp ./out/${SERVICE_NAME} /usr/local/bin

startd:
	docker-compose up -d --build