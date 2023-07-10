package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuttasakcom/go-fiber-simple/controllers"
)

func v1(a *fiber.App){
	route := a.Group("/api/v1")

	route.Post("/register", controllers.Register)
}
