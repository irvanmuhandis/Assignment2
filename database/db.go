package database

import (
	"assignment2/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

func StartDB() {
	psqInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(psqInfo), &gorm.Config{})
	fmt.Println(psqInfo)
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}
	db.Debug().AutoMigrate(model.Order{}, model.Item{})
}
