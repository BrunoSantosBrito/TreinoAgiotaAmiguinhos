package main

import (
	DB "TreinoAgiota/database"
	"TreinoAgiota/model"
	"TreinoAgiota/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
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

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "authorization", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com/"
		},
		MaxAge: 12 * time.Hour,
	}))
	routes.SetRoutes(router)
	router.Run()

}
