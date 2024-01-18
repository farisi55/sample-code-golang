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

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var age = 27
	rows, err := db.Query("SELECT id, name, grade FROM tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []student
	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func main() {

	sqlQuery()

}