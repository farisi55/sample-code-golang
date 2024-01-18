package database

import (
	"argoapi/entity"
	"log"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MYSQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	log.Println("Connection was successfully!!")
	return nil
}

//Migrate create/updates database table
func Migrate(table *entity.Person) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
