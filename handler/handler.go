package handler

import (
	"TreinoAgiota/model"
	"fmt"

	DB "TreinoAgiota/database"

	"github.com/gin-gonic/gin"
)

type H model.Va

func (h H) Post(c *gin.Context) {

	var agiota model.Agiota

	err := c.ShouldBindJSON(&agiota)
	if err != nil {
		return
	}

	err = DB.Conn.Create(&agiota).Error
	if err != nil {
		return
	}
	fmt.Println(agiota)
	c.JSON(200, gin.H{"response": agiota})

}
