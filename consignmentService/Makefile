build:
	GOOS=linux GOARCH=amd64 go build -tags netgo -o consignment-service
	docker build -t consignment-service .
run:
	docker run -p 50051:50051 consignment-service