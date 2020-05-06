package main

import (
	"coincover/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	FirstName   string `json:"first_name"`
	LastName	string `json:"last_name"`
	Country     string `json:"country"`
	Email	    string `json:"email"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/user", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":9090", router))

}

func createUser(w http.ResponseWriter, r *http.Request) {

	col := config.MS.Session.DB("db").C("users")

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

	//Insert job into MongoDB
	err = col.Insert(_u)
	if err != nil {
		panic(err)
	}

	//Convert job struct into json
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