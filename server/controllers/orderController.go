package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nao-18/golang-and-react-app/database"
	"github.com/nao-18/golang-and-react-app/models"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
