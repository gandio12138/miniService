package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/gandio12138/miniService/protobuf"
	"github.com/micro/go-micro"
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

func (s service) CreateConsignment(c context.Context, req *pb.CreateConsignmentReq, rsp *pb.CreateConsignmentRsp) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return errors.New(fmt.Sprintf("s.repo.Create() error: %v", err))
	}
	rsp.Created = true
	rsp.Consignment = consignment
	return nil
}

func (s service) GetConsignments(c context.Context, req *pb.GetConsignmentReq, rsp *pb.GetConsignmentRsp) error {
	resp := &pb.GetConsignmentRsp{
		Created:      true,
		Consignments: s.repo.GetAll(),
	}
	rsp = resp
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("go.micro.consignment.service"),
		micro.Version("latest"),
	)
	server.Init()
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo: &Repository{}})
	if err := server.Run(); err != nil {
		panic(err)
	}
}
