.PHONY: all test bench

test-v:
	go test ./... -v  -timeout=60000ms

test:
	go test ./...
	go test ./... -short -race
	go vet

bench: test
	go test ./... -test.run=NONE -test.bench=. -test.benchmem

get:
	go get

run: get
	go run main.go

run-race: get
	go run -race

linux: get
	GOOS=linux GOARCH=amd64 go build -v -o ./verisure-imap -ldflags '-s -w --extldflags "-static" ' ./main.go


all: bench run
