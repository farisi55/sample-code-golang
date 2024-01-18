package main

import (
	"fmt"
	"time"
	"log"
	"encoding/json"
	"net/http"


	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // turuna dari line 10
)

type GormModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

type User struct {
	GormModel
	Name  string `json:"name"`
	Email string `json:"email"`
}


func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "ormdb.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	var users []User
	db.Find(&users) // syarat mutlak jika berhubungan dengan query 
	fmt.Println("{}", users)
	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	db, err := gorm.Open("sqlite3", "ormdb.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	fmt.Println(name)
	fmt.Println(email)

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Succescly create")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "ormdb.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprint(w, "Successfully delete user")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "ormdb.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email

	db.Save(&user)

	fmt.Fprint(w, "Successfully update user")
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "ormdb.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	//Migrate the schema
	db.AutoMigrate(&User{})
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")
	initialMigration()
	handleRequests()

}