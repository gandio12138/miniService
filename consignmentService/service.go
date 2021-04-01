package main

import (
	"context"
	"errors"
	"fmt"
	consignmentPb "github.com/gandio12138/miniService/protobuf/consignment"
	vesselPb "github.com/gandio12138/miniService/protobuf/vessel"
	"github.com/micro/go-micro"
	"log"
)

type IRepository interface {
	Create(consignment *consignmentPb.CreateConsignmentReq) (*consignmentPb.Consignment, error)
	GetAll() []*consignmentPb.Consignment
}

type Repository struct {
	consignments []*consignmentPb.Consignment
}

func (repo *Repository) Create(consignment *consignmentPb.CreateConsignmentReq) (*consignmentPb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment.Consignment)
	return consignment.Consignment, nil
}

func (repo *Repository) GetAll() []*consignmentPb.Consignment {
	return repo.consignments
}

type service struct {
	repo         *Repository
	vesselClient vesselPb.VesselServiceClient
}

func (s service) CreateConsignment(_ context.Context, req *consignmentPb.CreateConsignmentReq, rsp *consignmentPb.CreateConsignmentRsp) error {
	vesselReq := &vesselPb.FindAvailableReq{
		Capacity:  int32(len(req.Consignment.Containers)),
		MaxWeight: req.Consignment.ConsignmentWeight,
	}
	findAvailableResp, err := s.vesselClient.FindAvailable(context.Background(), vesselReq)
	if err != nil {
		log.Fatalf("s.vesselClient.FindAvailable() error: %v\n", err)
	}
	log.Printf("found vessel: %s\n", findAvailableResp.Vessel.VesselName)
	req.Consignment.VesselId = findAvailableResp.Vessel.VesselId
	consignment, err := s.repo.Create(req)
	if err != nil {
		return errors.New(fmt.Sprintf("s.repo.Create() error: %v", err))
	}
	rsp.Created = true
	rsp.Consignment = consignment
	return nil
}

func (s service) GetConsignments(_ context.Context, _ *consignmentPb.GetConsignmentReq, rsp *consignmentPb.GetConsignmentRsp) error {
	resp := &consignmentPb.GetConsignmentRsp{
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
	vesselCli := vesselPb.NewVesselServiceClient("go.micro.service.vessel", server.Client())
	consignmentPb.RegisterShippingServiceHandler(server.Server(), &service{repo: &Repository{}, vesselClient: vesselCli})
	if err := server.Run(); err != nil {
		panic(err)
	}
}
