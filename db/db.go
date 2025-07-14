package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB * sql.DB


func InitDB() error {
 godotenv.Load()
 var err error
 
 dbHost := os.Getenv("DB_HOST")
 dbPort := os.Getenv("DB_PORT")
 dbUsername := os.Getenv("DB_USERNAME")
 dbName := os.Getenv("DB_NAME")
 dbPassword := os.Getenv("DB_PASSWORD")



 dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname="+ dbName + " port=" + dbPort + " sslmode=disable"
 DB, err = sql.Open("postgres", dsn)
 if err != nil {
  log.Println("error while starting the psql server: ", err)
  return err
 }
 if err = DB.Ping(); err != nil {
  log.Println("error making a test ping to the server: ", err)
  return err
 }
 log.Println("Database connected successfully")
 return nil
}