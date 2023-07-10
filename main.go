package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuttasakcom/go-fiber-simple/database"
	"github.com/yuttasakcom/go-fiber-simple/routes"
)

func main() {
	
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}