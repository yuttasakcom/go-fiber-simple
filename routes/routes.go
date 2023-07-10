package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuttasakcom/go-fiber-simple/controllers"
)

func Setup(app *fiber.App){
	app.Get("/", controllers.Hello)
}