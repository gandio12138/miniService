package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/gandio12138/miniService/protobuf"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"io/ioutil"
	"log"
	"os"
)

func parseFile(filePath string) (*pb.Consignment, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("ioutil.ReadFile() error: %v\n", err)
	}
	consignment := &pb.Consignment{}
	err = json.Unmarshal(fileData, consignment)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("json.Unmarshal() error: %v\n", err))
	}
	return consignment, nil
}

func main() {
	if err := cmd.Init(); err != nil {
		log.Fatalf("cmd.Init error: %v\n", err)
	}
	cli := pb.NewShippingServiceClient("go.micro.consignment.service", client.DefaultClient)
	infoFile := "consignment.json"
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("%v", err)
	}
	respSingle, err := cli.CreateConsignment(context.Background(), &pb.CreateConsignmentReq{Consignment: consignment})
	if err != nil {
		log.Fatalf("client.CreateConsignment() error: %v\n", err)
	}
	fmt.Printf("client.CreateConsignment resp: %v\n", respSingle)
	respAll, err := cli.GetConsignments(context.Background(), &pb.GetConsignmentReq{})
	if err != nil {
		log.Fatalf("client GetConsignment() error: %v\n", err)
	}
	fmt.Printf("client.GetConsignment() respAll: %v", respAll)
}
