BINARY := consignment-client

all: build run

build: deps
	@echo "=========> go build"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o $(BINARY)
	@echo "=========> docker build"
	docker build -t consignment-client .
deps:
	@echo "=========> go get"
	@go get -u github.com/asim/go-micro/v3
run:
	@echo "=========> docker run"
	docker run -e MICRO_REGISTER=mdns $(BINARY)

.PHONY: build deps run