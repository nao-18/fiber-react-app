package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nao-18/golang-and-react-app/database"
	"github.com/nao-18/golang-and-react-app/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
