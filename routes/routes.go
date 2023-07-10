package routes

import "github.com/gofiber/fiber/v2"

func Setup(a *fiber.App) {
	v1(a)
	SwaggerRoute(a)
}

