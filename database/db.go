package database

import (
	"log"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"go-heartbeat-job/models"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "root"
	dbPort   = "5432"
	dbname   = "climate"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	fmt.Println("successfully connected to database")
	db.AutoMigrate(&models.Weathers{})
}

func GetDB() *gorm.DB {
	return db
}