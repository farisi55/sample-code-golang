package main

import (
	"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	 // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    //db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlExec(){
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("insert into tb_student value (?,?,?,?)")
	stmt.Exec("G001", "Galand", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert success!")
}

func main () {
	sqlExec()
}