/*
 *     Copyright ATechnoHazard 2019  <amolele@gmail.com>.
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

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

	if _, ok := os.LookupEnv("HEROKU"); !ok {
		err := godotenv.Load() //Load .env file
		if err != nil {
			log.Fatalln(err)
		}
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")

	if _, ok := os.LookupEnv("HEROKU"); ok {

	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=require password=%s", dbHost, username, dbName,
		password) //Build the connection string

	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	log.Println("Auto-migrating db")
	db.Debug().AutoMigrate(&Account{}, &Notes{}) //Database migration
	log.Println("Auto-migration complete")
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
