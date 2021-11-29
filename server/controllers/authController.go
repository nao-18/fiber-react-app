package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nao-18/golang-and-react-app/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "john",
	}

	return c.JSON(user)
}
