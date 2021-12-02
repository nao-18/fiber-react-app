package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nao-18/golang-and-react-app/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "john",
	}

	user.LastName = "Doe"
	user.Email = "example@gmail.com"

	return c.JSON(user)
}
