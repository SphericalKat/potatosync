package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB //database

func init() {

	err := godotenv.Load() //Load .env file
	if err != nil {
		log.Fatalln(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=require password=%s", dbHost, username, dbName,
		password) //Build the connection string

	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	log.Println("Auto-migrating db")
	db.AutoMigrate(&Account{}) //Database migration
	log.Println("Auto-migration complete")
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
