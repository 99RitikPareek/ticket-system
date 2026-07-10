package config

import (
	"log"

	"ticket-system/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var err error

	DB, err = gorm.Open(sqlite.Open("ticket.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Ticket{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected Successfully")
}