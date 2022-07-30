package main

import (
	DB "TreinoAgiota/database"
	h "TreinoAgiota/handler"
	"TreinoAgiota/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB.DbConnection()

	DB.Conn.AutoMigrate(&model.Agiota{}, &model.Address{})

	batata := h.H{}

	app := gin.New()
	app.POST("/pojetinhocria", batata.Post)

	app.Run(":8080")

}
