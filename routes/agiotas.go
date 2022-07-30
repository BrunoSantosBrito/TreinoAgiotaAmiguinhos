package routes

import (
	"TreinoAgiota/handler"

	"github.com/gin-gonic/gin"
)

func SetAgiotaRoutes(r *gin.Engine) {
	agiota := r.Group("/agiota")
	controller := handler.AgiotaController{}

	{
		agiota.POST("/", controller.Post)
	}
}
