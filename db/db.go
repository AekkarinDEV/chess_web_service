package db

import (
	"log"
	"os"

	"github.com/AekkarinDEV/chess_web_service/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB * gorm.DB


func InitDB() error {
 godotenv.Load()
 var err error
 
 dbHost := os.Getenv("DB_HOST")
 dbPort := os.Getenv("DB_PORT")
 dbUsername := os.Getenv("DB_USERNAME")
 dbName := os.Getenv("DB_NAME")
 dbPassword := os.Getenv("DB_PASSWORD")



 dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname="+ dbName + " port=" + dbPort + " sslmode=disable"
 DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
 if err != nil {
  log.Println("error while starting the psql server: ", err)
  return err
 }

 DB.AutoMigrate(&models.User{})
 log.Println("database migrate completely")
 log.Println("Database connected successfully")
 return nil
}