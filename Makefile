.PHONY: build run stop

test: go-get
	go test ./...

lint: go-get
	golangci-lint run

go-get:
	go mod tidy

run:
	docker-compose up -d

stop:
	docker-compose stop
