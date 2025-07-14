package routes

import (
	"github.com/AekkarinDEV/chess_web_service/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(router fiber.Router){
	router.Post("/sign_up", controllers.SignUp)
	router.Post("/sign_in", controllers.SignIn)
}