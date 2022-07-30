package handler

import (
	"TreinoAgiota/model"

	DB "TreinoAgiota/database"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CapangaController struct{}

func (h *CapangaController) Post(c *gin.Context) {

	// var address model.Address
	var capanga model.Capanga

	err := c.ShouldBindBodyWith(&capanga, binding.JSON) // verifica JSON Capanga
	if err != nil {
		c.JSON(500, gin.H{"error: ": err.Error()})
		return
	}

	// err = c.ShouldBindBodyWith(&address, binding.JSON) // verifica JSON Address
	// if err != nil {
	// 	c.JSON(500, gin.H{"error: ": err.Error()})
	// 	return
	// }

	err = DB.Conn.Create(&capanga).Error
	if err != nil {
		c.JSON(500, gin.H{"error: ": err.Error()})
		return
	}
	c.JSON(200, gin.H{"response": capanga})

}
