package model

import (
	DB "TreinoAgiota/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Capanga struct {
	DB.Model

	Nome   string `json:"name" binding:"-"`
	Funcao string `json:"funcao" binding:"-"`
}

func (a *Capanga) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()

	return
}
