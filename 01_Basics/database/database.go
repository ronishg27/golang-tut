package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	DB = setupDatabase()
}

func setupDatabase() *gorm.DB {
	dbUrl := "host=localhost user=ronishg password=rons dbname=tododb port=5439 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbUrl),
		&gorm.Config{},
	)

	if err != nil {
		log.Println(err.Error())
		panic("failed to connect database")
	}
	log.Println("Connected to the database: tododb:5439")
	return db
}
