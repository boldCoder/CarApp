run:
	go run cmd/main.go

lint:
	golangci -lint ./...

build:
	go build -o main cmd/main.go
	