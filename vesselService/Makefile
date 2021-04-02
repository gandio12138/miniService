build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o vessel-service
	docker build -t vessel-service .
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service