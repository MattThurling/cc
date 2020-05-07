package main

import (
	"cc-api/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v10"
	"io/ioutil"
	"log"
	"net/http"
)

// User holds the user information and the mappings from json
type User struct {
	FirstName   string `json:"first_name" validate:"required,max=50"`
	LastName	string `json:"last_name" validate:"required,max=50"`
	Country     string `json:"country" validate:"required,min=2,max=50"`
	Email	    string `json:"email" validate:"required,email"`
}

// Program entrypoint
func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/user", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}

// createUser handles the POST request, saves the data into a struct and then writes to the database
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