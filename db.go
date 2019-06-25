package gql_company

import (
	"github.com/Shahid4183/gql-company/models"

	// gorm - it is an ORM library to work with SQL databases in go
	"github.com/jinzhu/gorm"

	// we are using mysql, thats why we will import mysql drivers
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// ConnectToDatabase - connects to mysql server running on localhost:3306
func ConnectToDatabase() error {
	// user = root
	// password = root
	// host = localhost/127.0.0.1
	// port = 3306
	// database = company
	// connection url = user:password@tcp(host:port)/database
	d, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/company")
	if err != nil {
		return err
	}
	db = d
	return nil
}

// GetDBInstance - gets database instance
func GetDBInstance() *gorm.DB {
	return db
}

// AutoMigrate - this function auto migrates the database tables
func AutoMigrate() {
	db.AutoMigrate(
		&models.Employee{},
		&models.Department{},
	)
}
