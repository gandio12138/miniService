package main

import (
	"context"
	"errors"
	"github.com/gandio12138/miniService/protobuf/vessel"
	"github.com/micro/go-micro"
	"log"
)

type Repository interface {
	FindAvailable(*vessel.FindAvailableReq) (*vessel.Vessel, error)
}

type VesselRepository struct {
	vessels []*vessel.Vessel
}

func (repo *VesselRepository) FindAvailable(req *vessel.FindAvailableReq) (*vessel.Vessel, error) {
	for _, v := range repo.vessels {
		if v.Capacity >= req.Capacity && v.MaxWeight >= req.MaxWeight {
			return v, nil
		}
	}
	return nil, errors.New("No vessel  be use")
}

type service struct {
	repo Repository
}

func (s *service) FindAvailable(context context.Context, req *vessel.FindAvailableReq, rsp *vessel.FindAvailableRsp) error {
	panic("implement me")
}

func main() {
	vessels := []*vessel.Vessel{
		{VesselId: "vessel001", VesselName: "Boat McBoldface1", MaxWeight: 200000, Capacity: 500},
		{VesselId: "vessel002", VesselName: "Boat McBoldface2", MaxWeight: 2, Capacity: 5},
	}
	server := micro.NewService(
		micro.Name("go.micro.service.vessel"),
		micro.Version("latest"),
	)
	server.Init()
	vessel.RegisterVesselServiceHandler(server.Server(), &service{
		repo: &VesselRepository{vessels: vessels},
	})
	if err := server.Run(); err != nil {
		log.Fatalf("server.Run() error: %v\n", err)
	}
}
