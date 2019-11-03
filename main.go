package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	poepApi "github.com/sheepover96/poep_api/api"
	db "github.com/sheepover96/poep_api/db"
	//. "github.com/sheepover96/poep/models"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	return router
}

func main() {
	// get .env
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	// db connection
	connPath := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dbConnection, err := sql.Open("mysql", connPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mysql.")
	err = dbConnection.Ping()
	if err != nil {
		panic(err)
	}
	db.DBcon = dbConnection
	defer dbConnection.Close()

	router := setupRouter()
	poepApi.PoemRoutes(router)
	poepApi.PoemThemeRoutes(router)
	router.Run()

}
