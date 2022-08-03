package routes

import (
	"TreinoAgiota/handler"

	"github.com/gin-gonic/gin"
)

func SetAgiotaRoutes(r *gin.Engine) {
	agiota := r.Group("/agiota")
	controler := handler.AgiotaControler{}

	{
		agiota.POST("/", controler.Post)
		agiota.GET("/GETALL", controler.GETALL)
		agiota.GET("/:id", controler.GETBYID)
		agiota.GET("/GETBYNAME/:name", controler.GETBYNAME)
		agiota.PUT("/:id", controler.UPDATE)
		agiota.DELETE("/:id", controler.DELETE)
	}
}
