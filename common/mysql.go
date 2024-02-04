package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"fmt"
	"log"
)

func ConnectDB() (*gorm.DB) {
	user     := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	name     := os.Getenv("MYSQL_DATABASE")
	url      := os.Getenv("MYSQL_URL")
	dsn      := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, url, name)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Connected to database %s!", name)

	return db
}
