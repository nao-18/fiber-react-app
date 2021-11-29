package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=postgres port=5432 user=postgres password=postgres dbname=local_db sslmode=disable TimeZone=Asia/Tokyo"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	fmt.Println(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":9000")
}
