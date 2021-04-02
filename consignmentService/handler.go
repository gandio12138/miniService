package main

import (
	"context"
	"errors"
	"fmt"
	consignmentPb "github.com/gandio12138/miniService/protobuf/consignment"
	"gopkg.in/mgo.v2"
	"log"
)
import vesselPb "github.com/gandio12138/miniService/protobuf/vessel"

type handler struct {
	session      *mgo.Session
	vesselClient vesselPb.VesselService
}

func (h *handler) CreateConsignment(ctx context.Context, req *consignmentPb.CreateConsignmentReq, rsp *consignmentPb.CreateConsignmentRsp) error {
	defer h.session.Close()
	vesselReq := &vesselPb.FindAvailableReq{
		Capacity:  int32(len(req.Consignment.Containers)),
		MaxWeight: req.Consignment.ConsignmentWeight,
	}
	resp, err := h.vesselClient.FindAvailable(ctx, vesselReq)
	if err != nil {
		return errors.New(fmt.Sprintf("in CreateConsignment call h.vesselClient.FindAvailable error: %v\n", err))
	}
	log.Printf("found vessel: %v\n", resp.Vessel.VesselName)
	req.Consignment.VesselId = resp.Vessel.VesselId
	if err = h.GetRepo().Create(req); err != nil {
		return errors.New(fmt.Sprintf("in CreateConsignment call h.GetRepo().Create() error: %v\n", err))
	}
	rsp = &consignmentPb.CreateConsignmentRsp{
		Created:     true,
		Consignment: req.Consignment,
	}
	return nil
}

func (h *handler) GetConsignments(ctx context.Context, req *consignmentPb.GetConsignmentReq, rsp *consignmentPb.GetConsignmentRsp) error {
	defer h.GetRepo().Close()
	cons, err := h.GetRepo().GetAll()
	if err != nil {
		return errors.New(fmt.Sprintf("in GetConsignments call h.GetRepo().GetAll() error: %v\n", err))
	}
	rsp.Consignments = cons
	return nil
}

func (h *handler) GetRepo() Repository {
	return &ConsignmentServiceRepository{sess: h.session.Clone()}
}
