BINARY := consignment-service

all: build run

build: deps
	@echo "========> go build"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o $(BINARY)
	docker build -t consignment-service .

deps:
	@go get -u github.com/asim/go-micro/v3
run:
	docker run -p 50051:50051 -e MICRO_ADDRESS=:50051 -e MICRO_REGISTER=mdns $(BINARY)

.PHONY: build deps run