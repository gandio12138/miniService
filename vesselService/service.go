package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/gandio12138/miniService/protobuf/vessel"
	"log"
	"os"
)

const (
	DbName        = "vessel"
	ConCollection = "consignments"
	DefaultHost   = "localhost:27017"
)

func createDummyData(repo Repository) error {
	defer repo.Close()
	vessels := []*vessel.Vessel{
		{VesselId: "vessel001", VesselName: "Boat McBoldface1", MaxWeight: 200000, Capacity: 500},
		{VesselId: "vessel002", VesselName: "Boat McBoldface2", MaxWeight: 2, Capacity: 5},
	}
	for _, vls := range vessels {
		if err := repo.Create(vls); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DefaultHost
	}
	sess, err := CreateMongoDBSession(dbHost)
	if err != nil {
		log.Fatalf("call CreateMongoDBSession() error: %v\n", err)
	}
	defer sess.Close()

	repo := &VesselServiceRepository{sess: sess}
	if err = createDummyData(repo); err != nil {
		log.Fatalf("call createDummyData error: %v\n", err)
	}
	server := micro.NewService(
		micro.Name("go.micro.service.vessel"),
		micro.Version("latest"),
	)
	server.Init()
	if err = vessel.RegisterVesselServiceHandler(server.Server(), &handler{sess: sess}); err != nil {
		log.Fatalf("vessel.RegisterVesselServiceHandler() error:%v\n", err)
	}
	if err := server.Run(); err != nil {
		log.Fatalf("server.Run() error: %v\n", err)
	}
}
