package routes

import "github.com/gofiber/fiber/v2"

func AuthRouter(router fiber.Router){
	router.Get("/", func(c * fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"messege": "hello from auth route",
		})
	})
}