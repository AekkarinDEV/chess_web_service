package routes

import "github.com/gofiber/fiber/v2"

func AppRouter(router fiber.Router){
	// auth router
	router.Route("/auth", AuthRouter)
}