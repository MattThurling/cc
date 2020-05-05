package main

import (
	"coincover/config"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	First   string `json:"first"`
	Last 	string `json:"last"`
	Country     string `json:"country"`
	Email      	string `json:"email"`
}

type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Salary      string `json:"salary"`
}


func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/jobs", jobsGetHandler).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":3333", router))

}

func jobsGetHandler(w http.ResponseWriter, r *http.Request) {

	col := config.Session.DB("db").C("jobs")

	results := []Job{}
	col.Find(bson.M{"title": bson.RegEx{"", ""}}).All(&results)
	jsonString, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(jsonString))

}

func createUser(w http.ResponseWriter, r *http.Request) {

	col := config.Session.DB("db").C("users")

	fmt.Print(col)
	//Retrieve body from http request
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	//Save data into User struct
	var _user User
	err = json.Unmarshal(b, &_user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Insert user into MongoDB
	err = col.Insert(_user)
	if err != nil {
		panic(err)
	}

	//Convert user struct into json
	jsonString, err := json.Marshal(_user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Set content-type http header
	w.Header().Set("content-type", "application/json")

	//Send back data as response
	w.Write(jsonString)

}
