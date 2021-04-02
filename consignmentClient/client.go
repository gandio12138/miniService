package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/cmd"
	consignmentPb "github.com/gandio12138/miniService/protobuf/consignment"
	"io/ioutil"
	"log"
	"os"
)

func parseFile(filePath string) (*consignmentPb.Consignment, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("ioutil.ReadFile() error: %v\n", err)
	}
	consignment := &consignmentPb.Consignment{}
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
	cli := consignmentPb.NewShippingService("go.micro.consignment.service", client.DefaultClient)
	infoFile := "consignment.json"
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("%v", err)
	}
	respSingle, err := cli.CreateConsignment(context.Background(), &consignmentPb.CreateConsignmentReq{Consignment: consignment})
	if err != nil {
		log.Fatalf("client.CreateConsignment() error: %v\n", err)
	}
	fmt.Printf("client.CreateConsignment resp: %v\n", respSingle)
	respAll, err := cli.GetConsignments(context.Background(), &consignmentPb.GetConsignmentReq{})
	if err != nil {
		log.Fatalf("client GetConsignment() error: %v\n", err)
	}
	fmt.Printf("client.GetConsignment() respAll: %v", respAll)
}
