package main

import (
	pb "github.com/gandio12138/miniService/protobuf"
	"log"
	"net"
)

type IRepository interface {
	Create(consignment *pb.CreateConsignmentReq) (*pb.CreateConsignmentRsp, error)
}

type Repository struct {
	consignments []*
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("net.Listen error: %v\n", err)
	}

}
