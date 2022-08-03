package handler

import (
	"TreinoAgiota/model"
	"fmt"
	"net/http"

	DB "TreinoAgiota/database"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AgiotaControler struct{}

func (ac *AgiotaControler) Post(c *gin.Context) {

	var agiota model.Agiota

	err := c.ShouldBindBodyWith(&agiota, binding.JSON)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	err = DB.Conn.Create(&agiota).Error
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	fmt.Println(agiota)
	c.JSON(200, gin.H{"response": agiota})

}

//function get all
func (ac *AgiotaControler) GETALL(c *gin.Context) {

	var agiota []model.Agiota

	err := DB.Conn.Joins("AddressAgiota").Joins("Capanga").Find(&agiota).Error
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	fmt.Println(agiota)
	c.JSON(200, gin.H{"response": agiota})
}

func (ac *AgiotaControler) GETBYID(c *gin.Context) {

	id := c.Param("id")

	var agiota model.Agiota

	query := DB.Conn.Joins("AddressAgiota").Joins("Capanga").Where("agiota.id = ?", id).First(&agiota)

	err := query.Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, agiota)
}

func (ac *AgiotaControler) GETBYNAME(c *gin.Context) {

	name := c.Param("name")

	var agiota model.Agiota

	query := DB.Conn.Joins("AddressAgiota").Joins("Capanga").Where("agiota.name = ?", name).First(&agiota)

	err := query.Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, agiota)
}

func (ac *AgiotaControler) UPDATE(c *gin.Context) {

	id := c.Param("id")

	var agiota model.Agiota

	err := c.ShouldBindJSON(&agiota)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	agiota.ID, err = uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = DB.Conn.Session(&gorm.Session{SkipHooks: true}).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&agiota).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"status": "Success!", "agiota": agiota})
}

func (ac *AgiotaControler) DELETE(c *gin.Context) {

	id := c.Param("id")

	err := DB.Conn.Delete(&model.Agiota{}, &id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success!"})
}
