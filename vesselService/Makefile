build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o vessel-service
	docker build -t vessel-service .
run:
	docker run -p 50052:50052 -e MICRO_REGISTER=mdns vessel-service