package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nao-18/golang-and-react-app/database"
	"github.com/nao-18/golang-and-react-app/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)
	app.Listen(":9000")
}
