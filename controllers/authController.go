package controllers

import (
	"log"
	"strings"

	"github.com/AekkarinDEV/chess_web_service/db"
	"github.com/AekkarinDEV/chess_web_service/models"
	"github.com/AekkarinDEV/chess_web_service/utils"
	"github.com/gofiber/fiber/v2"
)

type signUpRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}


func SignUp(c *fiber.Ctx) error{
	var req signUpRequest

	// request validate
	if err := c.BodyParser(&req); err != nil {
		log.Println(req.Username)
		return c.Status(400).JSON(fiber.Map{
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
				return c.Status(409).JSON(fiber.Map{
					"error": "username used",
				})
			}
		}
		return c.Status(500).JSON(fiber.Map{
			"error": dbRespond.Error.Error(),
		})
	}

	//ok 
	return c.Status(200).JSON(fiber.Map{
		"messege" : "user created",
	})
}

type signInRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(c * fiber.Ctx) error {
	var req signUpRequest

	// request validate
	if err := c.BodyParser(&req); err != nil {
		log.Println(req.Username)
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request object",
		})
	}

	var user models.User
	res := db.DB.Where("username = ?", req.Username).First(&user)
	if res.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	if !utils.CompareWithHashedPassword(user.Password, req.Password) {
		return c.Status(400).JSON(fiber.Map{
			"error": "incorrect password",
		})
	}

	token, err := utils.GenerateJWTToken(user.Id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//ok
	return c.Status(200).JSON(fiber.Map{
		"username": user.Username,
		"access_token": token,
	})

}