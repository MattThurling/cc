package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// User holds the user information and the mappings from json
type User struct {
	FirstName   string `json:"first_name" validate:"required,max=50"`
	LastName	string `json:"last_name" validate:"required,max=50"`
	Country     string `json:"country" validate:"required,min=2,max=50"`
	Email	    string `json:"email" validate:"required,email"`
}

// MongoStore stores the Mongo session
type MongoStore struct {
	session *mgo.Session
}

var mongoStore = MongoStore{}

func initMongo() (session *mgo.Session) {

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

	return

}


// Program entrypoint
func main() {

	//Create MongoDB session
	session := initMongo()
	mongoStore.session = session

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/user", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}

// createUser handles the POST request, saves the data into a struct and then writes to the database
func createUser(w http.ResponseWriter, r *http.Request) {

	col := mongoStore.session.DB(os.Getenv("DB_DATABASE")).C("users")

	//Retrieve body from http request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	//Save data into User struct
	var _u User
	err = json.Unmarshal(b, &_u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Validate input data
	v := validator.New()
	err = v.Struct(_u)
	if err != nil {
		// TODO: write custom production error messages - and maybe concatenate
		// For now, return first error's defualt message
		http.Error(w, err.(validator.ValidationErrors)[:1].Error(), 422)
		return
	}

	//Insert user into MongoDB
	err = col.Insert(_u)
	if err != nil {
		panic(err)
	}

	//Convert user struct into json
	jsonString, err := json.Marshal(_u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Set content-type http header
	w.Header().Set("content-type", "application/json")

	//Send back data as response
	w.Write(jsonString)

}