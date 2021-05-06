BINARY := vessel-service

all: build run

build: deps
	@echo "=========> go build"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o $(BINARY)
	docker build -t vessel-service .
deps:
	@go get -u github.com/asim/go-micro/v3
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns $(BINARY)

.PHONY: build deps run