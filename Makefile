.PHONY: build run stop build

build: go-get
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	go build -o main -i cmd/main.go

test: go-get
	go test ./...

lint: go-get
	golangci-lint run

go-get:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.30.0
	go mod tidy

run:
	docker-compose up -d

stop:
	docker-compose stop
