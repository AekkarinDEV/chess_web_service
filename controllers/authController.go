package controllers

import (
	"log"
	"strings"

	"github.com/AekkarinDEV/chess_web_service/db"
	"github.com/AekkarinDEV/chess_web_service/models"
	"github.com/AekkarinDEV/chess_web_service/utils"
	"github.com/gofiber/fiber/v2"
)

type signInRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(context *fiber.Ctx) error{
	var req signInRequest

	// request validate
	if err := context.BodyParser(&req); err != nil {
		log.Println(req.Username)
		return context.Status(400).JSON(fiber.Map{
			"error": "invalid request object",
		})
	}

	//queries
	newUser := models.User{
		Id: utils.GenerateRecordID(),
		Username: req.Username,
		Password: utils.GenerateHashedPassword(req.Password),
	}

	dbRespond := db.DB.Create(&newUser)

	//error handler
	if dbRespond.Error != nil {
		errMessege := dbRespond.Error.Error()
		if strings.Contains(errMessege,"duplicate"){
			//duplicate username
			if strings.Contains(errMessege,"username"){
				return context.Status(409).JSON(fiber.Map{
					"error": "username used",
				})
			}
		}
		return context.Status(500).JSON(fiber.Map{
			"error": dbRespond.Error.Error(),
		})
	}

	//ok 
	return context.Status(200).JSON(fiber.Map{
		"messege" : "user created",
	})
}