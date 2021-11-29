package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(postgres.Open("host=postgres port=5432 user=postgres password=postgres dbname=local_db sslmode=disable TimeZone=Asia/Tokyo"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
}
