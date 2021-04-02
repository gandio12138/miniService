package main

import (
	"github.com/asim/go-micro/v3"
	consignmentPb "github.com/gandio12138/miniService/protobuf/consignment"
	vesselPb "github.com/gandio12138/miniService/protobuf/vessel"
	"log"
	"os"
)

const (
	DefaultHost = "localhost:27017"
	DbName        = "consignment"
	ConCollection = "consignments"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DefaultHost
	}
	sess, err := CreateMongoDBSession(dbHost)
	if err != nil {
		log.Fatalf("call CreateMongoDBSession error: %v\n", err)
	}
	defer sess.Close()
	server := micro.NewService(
		micro.Name("go.micro.consignment.service"),
		micro.Version("latest"),
	)
	server.Init()
	vesselCli := vesselPb.NewVesselService("go.micro.service.vessel", server.Client())
	if err := consignmentPb.RegisterShippingServiceHandler(server.Server(), &handler{session: sess, vesselClient: vesselCli}); err != nil {
		log.Fatalf("consignmentPb.RegisterShippingServiceHandler error: %v\n", err)
	}
	if err := server.Run(); err != nil {
		panic(err)
	}
}
