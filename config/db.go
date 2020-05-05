package config

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

const (
	hosts      = "localhost:27017"
	username   = ""
	password   = ""
)

var Session *mgo.Session

func init() {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Username: username,
		Password: password,
	}

	Session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	err = Session.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	return

}