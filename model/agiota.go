package model

import (
	DB "TreinoAgiota/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Va struct{}

type Agiota struct {
	DB.Model
	Name          string    `json:"name" binding:"-" `
	Valor         string    `json:"valor" binding:"-" `
	AddressID     uuid.UUID `json:"address_id" binding:"-"`
	AddressAgiota Address   `json:"address_agiota"  binding:"-" gorm:"foreignKey:AddressID"`
}

func (a *Agiota) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()

	return
}
