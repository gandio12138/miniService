build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -o consignment-client
	docker build -t consignment-client .
run:
	docker run -e MICRO_REGISTER=mdns consignment-client