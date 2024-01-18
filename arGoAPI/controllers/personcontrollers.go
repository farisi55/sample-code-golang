package controllers

import (
	"argoapi/database"
	"argoapi/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// allow
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
}


//GetAllPerson get all data person
func GetAllPerson(w http.ResponseWriter, r *http.Request){
	enableCors(&w)

	var persons []entity.Person
	database.Connector.Find(&persons)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&persons)
}

//GetPersonById returns person with spesific ID
func GetPersonById(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	key := vars["id"]

	var persons entity.Person
	database.Connector.First(&persons, key)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&persons)
}

//CreatePerson 
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	requestBody, _ := ioutil.ReadAll(r.Body)
	var persons entity.Person
	json.Unmarshal(requestBody, &persons)
	database.Connector.Create(&persons)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&persons)
}

//UpadatePersonByID update person with respective ID
func UpadatePersonByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	requestBody, _ := ioutil.ReadAll(r.Body)
	var persons entity.Person
	json.Unmarshal(requestBody, &persons)
	database.Connector.Save(&persons)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&persons)
} 

//DeletePersonByID delete's person with specisfic ID
func DeletePersonByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	key := vars["id"]

	var persons entity.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&persons)
	w.WriteHeader(http.StatusNoContent)
}

