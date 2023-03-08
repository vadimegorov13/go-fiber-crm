package db

import (
	"log"

	"github.com/vadimegorov13/go-fiber-crm/pkg/common/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		// this will be printed in the terminal, confirming the connection to the database
		log.Println("The database is connected")
	}

	db.AutoMigrate(&models.Lead{})
	log.Println("Dtabase migrated!")

	return db
}
