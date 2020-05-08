// Package config sets up the connection to the Mongo database container
package config

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

// MongoStore stores the Mongo session
type MongoStore struct {
	Session *mgo.Session
}

var MS = MongoStore{}

func init() {

	info := &mgo.DialInfo{
		Addrs:    []string{os.Getenv("DB_HOST")},
		Timeout:  60 * time.Second,
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
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