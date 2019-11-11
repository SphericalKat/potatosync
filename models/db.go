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
	"log"
	"os"
	"time"

	u "github.com/ATechnoHazard/potatosync/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // import since we require the postgres dialect
	"github.com/joho/godotenv"
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
	dbSSL := os.Getenv("db_ssl")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, username, dbName,
		dbSSL, password) //Build the connection string

	conn, err := gorm.Open(dbType, dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	log.Println("Auto-migrating db")
	db.AutoMigrate(&Account{}, &Notes{}) //Database migration
	log.Println("Auto-migration complete")
}

//GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

func Stats() map[string]interface{} {
	userCount, notesCount, recentUsers := 0, 0, 0

	res := u.Message(true, "success")

	err := GetDB().Model(&Account{}).Count(&userCount).Error
	if err != nil {
		return u.Message(false, err.Error())
	}
	res["user_count"] = userCount

	err = GetDB().Model(&Notes{}).Count(&notesCount).Error
	if err != nil {
		return u.Message(false, err.Error())
	}
	res["notes_count"] = notesCount

	yesterday := time.Now().AddDate(0, 0, -1) // Get yesterday's time
	err = GetDB().Model(&Account{}).Where("created_at >= ?", yesterday).Count(&recentUsers).Error
	if err != nil {
		return u.Message(false, err.Error())
	}
	res["recent_users"] = recentUsers

	return res
}
