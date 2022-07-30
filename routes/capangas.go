package routes

import (
	"TreinoAgiota/handler"

	"github.com/gin-gonic/gin"
)

func SetCapangasRoutes(r *gin.Engine) {
	capanga := r.Group("/capanga")
	controller := handler.CapangaController{}

	{
		capanga.POST("/", controller.Post)
	}
}
