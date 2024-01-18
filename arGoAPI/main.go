package main

import (
	"argoapi/database"
	"argoapi/controllers"
	"argoapi/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //required for mysql dialects
)

//intialisai database
func arInitDB() {
	config :=
	database.Config{
		ServerName: "localhost:3306",
		User: "root",
		Password: "",
		DB: "argoapidb",
	}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&entity.Person{})
}

func intialiseHandler(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonById).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.GetPersonById).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.GetPersonById).Methods("DELETE")
}

func main() {
	arInitDB()

	//simple http server
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)

	intialiseHandler(router)

	log.Fatal(http.ListenAndServe(":8090", router))
	//end simple http server

}