package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/gandio12138/miniService/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

type IRepository interface {
	Create(consignment *pb.CreateConsignmentReq) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.CreateConsignmentReq) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment.Consignment)
	return consignment.Consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo *Repository
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetConsignmentReq) (*pb.GetConsignmentRsp, error) {
	resp := &pb.GetConsignmentRsp{
		Created:      true,
		Consignments: s.repo.GetAll(),
	}
	return resp, nil
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.CreateConsignmentReq) (*pb.CreateConsignmentRsp, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("s.repo.Create() error: %v", err))
	}
	return &pb.CreateConsignmentRsp{
		Created:     true,
		Consignment: consignment,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("net.Listen error: %v\n", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterShippingServiceServer(grpcServer, &service{repo: &Repository{}})
	fmt.Println("start grpc service......")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpcService.Server() error: %v\n", err)
	}
}
