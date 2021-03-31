build:
	GOOS=linux GOARCH=amd64 go build -tags netgo -o consignment-client
	docker build -t consignment-client .
run:
	docker run consignment-client