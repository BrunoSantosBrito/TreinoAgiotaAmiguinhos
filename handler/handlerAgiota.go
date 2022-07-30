package handler

import (
	"TreinoAgiota/model"
	"fmt"

	DB "TreinoAgiota/database"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AgiotaController struct{}

func (h *AgiotaController) Post(c *gin.Context) {

	var agiota model.Agiota
	// var address model.Address

	err := c.ShouldBindBodyWith(&agiota, binding.JSON) // verifica JSON Agiota
	if err != nil {
		c.JSON(500, gin.H{"error: ": err.Error()})
		return
	}
	fmt.Println(agiota)
	err = c.ShouldBindBodyWith(&agiota.AddressAgiota, binding.JSON) // verifica JSON Address
	if err != nil {
		c.JSON(500, gin.H{"error: ": err.Error()})
		return
	}

	err = DB.Conn.Create(&agiota).Error
	if err != nil {
		c.JSON(500, gin.H{"error: ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": agiota})

}
