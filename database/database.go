package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"webserser/utils"
)

var DB *gorm.DB

func GetDatabase(dbName string) *gorm.DB {

	// initialize some variables
	// for the MySQL data source
	var (
		databaseUser     string = utils.GetValue("DB_USER")
		databasePassword string = utils.GetValue("DB_PASSWORD")
		databaseHost     string = utils.GetValue("DB_HOST")
		databasePort     string = utils.GetValue("DB_PORT")
		databaseName     string = dbName
	)

	// declare the data source for MySQL
	var dataSource = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", databaseHost, databaseUser, databasePassword, databaseName, databasePort)
	// create a variable to store an error
	var err error

	// create a connection to the database
	DB, err = gorm.Open(postgres.Open(dataSource), &gorm.Config{})

	// if connection fails, print out the errors
	if err != nil {
		panic(err.Error())
	}

	return DB
}
