package repository

import (
	"smart-kost-backend/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(inputModel *model.UserList) (*model.UserList, error)
	GetUserByUsername(inputModel *model.UserList) (*model.UserList, error)
	UpdateIsOnline(inputModel *model.UserList) (*model.UserList, error)
	UpdateIsSOS(inputModel *model.UserList) (*model.UserList, error)
	UpdateUserOffline()
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r userRepo) WithTx(tx *gorm.DB) UserRepository {
	return &userRepo{
		db: tx,
	}
}

func (r *userRepo) CreateUser(inputModel *model.UserList) (*model.UserList, error) {
	res := r.db.Raw("INSERT INTO user_list (username , pass , is_online, is_sos) VALUES (?,?,?,?) RETURNING *;", inputModel.Username, inputModel.Pass, "0", "0").Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}

func (r *userRepo) GetUserByUsername(inputModel *model.UserList) (*model.UserList, error) {

	res := r.db.Where("username = ?", inputModel.Username).First(&model.UserList{}).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}

func (r *userRepo) UpdateIsOnline(inputModel *model.UserList) (*model.UserList, error) {

	res := r.db.Raw("UPDATE user_list SET is_online = ?, updated_at = now()+ interval '7 hour' WHERE user_id = ? RETURNING *", string(inputModel.IsOnline), inputModel.UserId).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}

	return inputModel, nil
}

func (r *userRepo) UpdateIsSOS(inputModel *model.UserList) (*model.UserList, error) {

	res := r.db.Raw("UPDATE user_list SET is_sos = ?, updated_at = now()+ interval '7 hour' WHERE user_id = ? RETURNING *", string(inputModel.IsSOS), inputModel.UserId).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}

	return inputModel, nil
}

func (r *userRepo) UpdateUserOffline() {
	var modelDb model.UserList
	res := r.db.Raw("UPDATE user_list SET is_online = '0' WHERE is_online = '1' AND extract(epoch from (now()+ interval '7 hour') - updated_at) / 60 > 5").Scan(&modelDb)
	if res.Error != nil {
		println(res.Error)
	}
}
