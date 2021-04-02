package main

import (
	consignmentPb "github.com/gandio12138/miniService/protobuf/consignment"
	"gopkg.in/mgo.v2"
)

type Repository interface {
	Create(consignment *consignmentPb.CreateConsignmentReq) error
	GetAll() ([]*consignmentPb.Consignment, error)
	Close()
}

type ConsignmentServiceRepository struct {
	sess *mgo.Session
}

func (repo *ConsignmentServiceRepository) collection() *mgo.Collection {
	return repo.sess.DB(DbName).C(ConCollection)
}

func (repo *ConsignmentServiceRepository) Create(consignment *consignmentPb.CreateConsignmentReq) error {
	return repo.collection().Insert(consignment.Consignment)
}
func (repo *ConsignmentServiceRepository) GetAll() ([]*consignmentPb.Consignment, error) {
	var cons []*consignmentPb.Consignment
	err := repo.collection().Find(nil).All(&cons)
	return cons, err
}
func (repo *ConsignmentServiceRepository) Close() {
	repo.sess.Close()
}
