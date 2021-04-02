package main

import (
	"errors"
	"github.com/gandio12138/miniService/protobuf/vessel"
	"gopkg.in/mgo.v2"
	"log"
)

type Repository interface {
	FindAvailable(*vessel.FindAvailableReq) (*vessel.Vessel, error)
	Create(*vessel.Vessel) error
	Close()
}

type VesselServiceRepository struct {
	sess *mgo.Session
}

func (repo *VesselServiceRepository) collection() *mgo.Collection {
	return repo.sess.DB(DbName).C(ConCollection)
}

func (repo *VesselServiceRepository) FindAvailable(req *vessel.FindAvailableReq) (*vessel.Vessel, error) {
	var vessels []*vessel.Vessel
	if err := repo.collection().Find(nil).All(&vessels); err != nil {
		log.Fatalf("in FindAvailable call repo.collection().Find error: %v\n", err)
	}

	for _, v := range vessels {
		if v.Capacity >= req.Capacity && v.MaxWeight >= req.MaxWeight {
			return v, nil
		}
	}
	return nil, errors.New("No vessel  be use")
}

func (repo *VesselServiceRepository) Create(req *vessel.Vessel) error {
	return repo.collection().Insert(req)
}

func (repo *VesselServiceRepository) Close() {
	repo.sess.Close()
}
