package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gandio12138/miniService/protobuf/vessel"
	"gopkg.in/mgo.v2"
)

type handler struct {
	sess *mgo.Session
}

func (h *handler) Create(_ context.Context, req *vessel.CreateReq, rsp *vessel.CreateRsp) error {
	err := h.GetRepo().Create(req.VesselInfo)
	if err != nil {
		return err
	}
	rsp.VesCreated = true
	rsp.VesselInfos = []*vessel.Vessel{req.VesselInfo}
	return nil
}

func (h *handler) GetRepo() Repository {
	return &VesselServiceRepository{sess: h.sess.Clone()}
}

func (h *handler) FindAvailable(_ context.Context, req *vessel.FindAvailableReq, rsp *vessel.FindAvailableRsp) error {
	vests, err := h.GetRepo().FindAvailable(req)
	if err != nil {
		return errors.New(fmt.Sprintf("in FindAvailable call h.GetRepo().FindAvailable() error: %v\n", err))
	}
	rsp.Vessel = vests
	return nil
}
