package database

import (
	"github.com/nao-18/golang-and-react-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(postgres.Open("host=postgres port=5432 user=postgres password=postgres dbname=local_db sslmode=disable TimeZone=Asia/Tokyo"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = database

	database.AutoMigrate(&models.User{})
}
