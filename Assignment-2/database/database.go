package database

import (
	"fmt"
	"log"

	"github.com/roihan12/assignment-2/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root123"
	dbName   = "hactiv8_order_go"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(&model.Order{}, &model.Item{})
}

func DB() *gorm.DB {
	return db
}
