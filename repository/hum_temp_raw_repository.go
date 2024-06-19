package repository

import (
	"smart-kost-backend/model"

	"gorm.io/gorm"
)

type HumTempRawRepository interface {
	CreateHumTempRawData(inputModel *model.HumTemRaw) (*model.HumTemRaw, error)
}

type humTempRawRepo struct {
	db *gorm.DB
}

func NewHumTempRawRepository(db *gorm.DB) HumTempRawRepository {
	return &humTempRawRepo{
		db: db,
	}
}

func (r humTempRawRepo) WithTx(tx *gorm.DB) HumTempRawRepository {
	return &humTempRawRepo{
		db: tx,
	}
}

func (r *humTempRawRepo) CreateHumTempRawData(inputModel *model.HumTemRaw) (*model.HumTemRaw, error) {

	if dbc := r.db.Create(inputModel).Scan(inputModel); dbc.Error != nil {
		return nil, dbc.Error
	}

	return inputModel, nil
}
