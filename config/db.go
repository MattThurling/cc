package config

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

const (
	username   = ""
	password   = ""
	database   = "db"
)

type MongoStore struct {
	Session *mgo.Session
}

var MS = MongoStore{}

func init() {

	info := &mgo.DialInfo{
		Addrs:    []string{os.Getenv("DB_HOST")},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	err = session.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	MS.Session = session

}