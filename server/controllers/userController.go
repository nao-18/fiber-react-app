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
