package main

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
)

func CreateMongoDBSession(host string) (*mgo.Session, error) {
	mgoSess, err := mgo.Dial(host)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("in CreateMongoDBSession mgo.Dial() error: %v\n", err))
	}
	mgoSess.SetMode(mgo.Monotonic, true)
	return mgoSess, nil
}
