package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nao-18/golang-and-react-app/database"
	"github.com/nao-18/golang-and-react-app/models"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)
	return c.JSON(user)
}
