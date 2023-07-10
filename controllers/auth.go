package controllers

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/yuttasakcom/go-fiber-simple/medels"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	user.FirstName = "Yuttasak"
	user.LastName = "Pannawat"
	user.Email = "yuttasakcom@gmail.com"
	return c.JSON(user)
}