package main

import (
	"log"

	"github.com/AekkarinDEV/chess_web_service/db"
	"github.com/AekkarinDEV/chess_web_service/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	err := db.InitDB()
	if err != nil{
		log.Print("error while connecting to database")
	}
	// err = db.MigrateDb()
	// if err != nil{
	// 	log.Print("error while migrating database")
	// }

	app.Get("/", func(c * fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"messege" : "hello world",
		})
	})

	//init routing
	app.Route("/api/v1/",routes.AppRouter)

	err = app.Listen(":8000")

	if(err != nil){
		log.Print("error on starting server:" ,err)
	}

}
