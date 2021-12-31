package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB = nil
var err error = nil

/**
* connect with data base with ..env file params
* just edit all data in ..env file
 */
func ConnectToDatabase() {
	if DB == nil {
		DB, err = gorm.Open("mysql", os.Getenv("DATABASE_USERNAME")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+":"+os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local&character_set_server=utf8mb4")
	}
	if err != nil {
		fmt.Println("Connect To Database Error:", err.Error())
		return
	}
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_DATABASE"))
	if os.Getenv("APP_ENV") == "local" {
		DB.LogMode(debug)
	}
}
