package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/yuttasakcom/go-fiber-simple/database"
	models "github.com/yuttasakcom/go-fiber-simple/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password not match",
		})
	}

	password,_ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := new(models.User)
	user.FirstName = data["first_name"]
	user.LastName = data["last_name"]
	user.Email = data["email"]
	user.Password = password

	database.DB.Create(&user)

	return c.JSON(user)
}