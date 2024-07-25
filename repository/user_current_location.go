package repository

import (
	"smart-kost-backend/model"

	"gorm.io/gorm"
)

type UserCurrentLocationRepo interface {
	CreateCurrentUserData(inputModel *model.UserCurrentLocation) (*model.UserCurrentLocation, error)
}

type userCurrentLocationRepo struct {
	db *gorm.DB
}

func NewUserCurrentLocationRepository(db *gorm.DB) UserCurrentLocationRepo {
	return &userCurrentLocationRepo{
		db: db,
	}
}

func (r userCurrentLocationRepo) WithTx(tx *gorm.DB) UserCurrentLocationRepo {
	return &userCurrentLocationRepo{
		db: tx,
	}
}

func (r *userCurrentLocationRepo) CreateCurrentUserData(inputModel *model.UserCurrentLocation) (*model.UserCurrentLocation, error) {

	res := r.db.Raw("INSERT INTO user_current_location (user_id , status_id , lat, long) VALUES (?,?,?,?) RETURNING *;", inputModel.UserId, inputModel.StatusId, inputModel.UserCurrentLocationLat, inputModel.UserCurrentLocationLong).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}

	return inputModel, nil

}
