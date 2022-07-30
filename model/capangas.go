package model

import (
	DB "TreinoAgiota/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Capanga struct {
	DB.Model
	NomeCapanga    string  `json:"nomeCapanga" binding:"required"`
	Especializacao string  `json:"especializacao" binding:"required"`
	Contato        string  `json:"contato" binding:"required"`
	TipoAbordagem  string  `json:"tipo de abordagem" binding:"-"`
	ValorServico   float64 `json:"valor do servico" binding:"required"`
	NaoFaco        string  `json:"Não trabalho com" binding:"-"`
}

func (a *Capanga) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
