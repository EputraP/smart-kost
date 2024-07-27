package repository

import (
	"smart-kost-backend/model"

	"gorm.io/gorm"
)

type UserCurrentLocationRepo interface {
	CreateCurrentUserData(inputModel *model.UserCurrentLocation) (*model.UserCurrentLocation, error)
	GetCurrentUserData() (*model.GetUserCurrentLocation, error)
	UpdateLatLong(inputModel *model.UserCurrentLocation) (*model.UserCurrentLocation, error)
	UpdateLocationStatus()
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

	res := r.db.Raw("INSERT INTO user_current_location (user_id , status_id , lat, long, prev_lat, prev_long) VALUES (?,?,?,?,?,?) RETURNING *;", inputModel.UserId, inputModel.StatusId, inputModel.UserCurrentLocationLat, inputModel.UserCurrentLocationLong, nil, nil).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}

	return inputModel, nil
}

func (r *userCurrentLocationRepo) GetCurrentUserData() (*model.GetUserCurrentLocation, error) {
	var dbResultModel *model.GetUserCurrentLocation

	sqlScript := "select ul.username, ul.is_online, ul.is_sos, sl.status_name, ucl.lat, ucl.long from user_current_location ucl left join user_list ul  on ucl.user_id  = ul.user_id  left join status_list sl on ucl.status_id = sl.status_id"

	res := r.db.Raw(sqlScript).Scan(&dbResultModel)
	if res.Error != nil {
		return nil, res.Error
	}

	return dbResultModel, nil
}

func (r *userCurrentLocationRepo) UpdateLatLong(inputModel *model.UserCurrentLocation) (*model.UserCurrentLocation, error) {

	res := r.db.Raw("UPDATE user_current_location SET lat = ?, long = ?, updated_at = now()+ interval '7 hour' WHERE user_id = ? RETURNING *", inputModel.UserCurrentLocationLat, inputModel.UserCurrentLocationLong, inputModel.UserId).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}

	return inputModel, nil
}

func (r *userCurrentLocationRepo) UpdateLocationStatus() {
	var modelDb model.UserCurrentLocation
	res := r.db.Raw("UPDATE user_current_location SET status_id  = case WHEN (extract(epoch from (now()+ interval '7 hour') - updated_at) / 60 < 5) and (prev_lat is not null or prev_long is not null) THEN 1 else 2 end").Scan(modelDb)
	if res.Error != nil {
		println(res.Error)
	}

}
