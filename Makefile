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
	GOOS=linux GOARCH=amd64 go build -v -o ./sonos-domotics -ldflags '-s -w --extldflags "-static" ' ./main.go ./api.go


docker: get
	docker build -t sonos-go:1.0 . -f Dockerfile

all: bench run
