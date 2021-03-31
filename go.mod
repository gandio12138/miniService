module github.com/gandio12138/miniService

go 1.14

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro v1.18.0
	golang.org/x/net v0.0.0-20210330230544-e57232859fb2
	google.golang.org/genproto v0.0.0-20210330181207-2295ebbda0c6 // indirect
)

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
